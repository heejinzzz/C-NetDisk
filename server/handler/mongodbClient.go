package handler

import (
	pb "C-NetDisk/server/proto"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	pt "path"
)

const (
	mongodbServerURI    = "mongodb://10.96.2.10:27017"
	dbName              = "CNetDisk"
	usersCollectionName = "users"
	itemsCollectionName = "items"
)

var db *mongo.Database

func init() {
	fmt.Println("C-NetDisk Server Start.")
	fmt.Println("Connecting MongoDB...")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbServerURI))
	if err != nil {
		log.Fatalln("[FATAL]Failed to connect MongoDB. C-NetDisk Server Exit.")
	}
	fmt.Println("Connect MongoDB succeed.\nService Start. Waiting For Connection...")
	db = client.Database(dbName)
}

func isUsernameValid(username string) bool {
	if len(username) < UsernameMinLength || len(username) > UsernameMaxLength {
		return false
	}
	for _, v := range []byte(username) {
		if ForbiddenChar[v] {
			return false
		}
	}
	return true
}

func isUserExist(username string) bool {
	collection := db.Collection(usersCollectionName)
	cur, err := collection.Find(context.Background(), bson.M{"name": username})
	if err != nil {
		log.Fatalln(err)
	}
	if cur.Next(context.Background()) {
		return true
	}
	return false
}

func createNewUser(username string, password string) error {
	secret := AesEncrypt(password)
	_, err := db.Collection(usersCollectionName).InsertOne(context.Background(), bson.M{"name": username, "password": secret})
	if err != nil {
		return err
	}

	path := StorageRootDirectory + "/" + username
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	item := pb.ItemInfo{
		Name:     username,
		Type:     pb.ItemType_File,
		Children: []*pb.ItemBasicInfo{},
		Size:     0,
		Path:     path,
	}
	_, err = db.Collection(itemsCollectionName).InsertOne(context.Background(), item)
	return err
}

func checkUserPassword(username string, password string) bool {
	if !isUserExist(username) {
		return false
	}
	if password == AesDecrypt(getUserPassword(username)) {
		return true
	}
	return false
}

func getUserPassword(username string) string {
	res := db.Collection(usersCollectionName).FindOne(context.Background(), bson.M{"name": username})
	var mp map[string]string
	err := res.Decode(&mp)
	if err != nil {
		log.Fatalln(err)
	}
	return mp["password"]
}

func getItemInfoFromMongo(name string) (*pb.ItemInfo, error) {
	res := db.Collection(itemsCollectionName).FindOne(context.Background(), bson.M{"name": name})
	var item pb.ItemInfo
	err := res.Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func getItemPath(name string) string {
	info, err := getItemInfoFromMongo(name)
	if err != nil {
		log.Fatalln(err)
	}
	return info.Path
}

func createItemInMongo(createInfo *pb.ItemInfo) (*pb.ItemInfo, error) {
	if len(createInfo.Name) < FilenameMinLength {
		return nil, errors.New("the file name is too short")
	}
	if len(createInfo.Name) > FilenameMaxLength {
		return nil, errors.New("the file name is too long")
	}
	lastDirInfo, err := getItemInfoFromMongo(createInfo.Path)
	if err != nil {
		return nil, err
	}
	path := lastDirInfo.Path + "/" + base64Encode(createInfo.Name)
	name := createInfo.Path + "/" + createInfo.Name
	itemInfo := pb.ItemInfo{
		Name:     name,
		Type:     createInfo.Type,
		Size:     createInfo.Size,
		Path:     path,
		Children: createInfo.Children,
	}
	_, err = db.Collection(itemsCollectionName).InsertOne(context.Background(), itemInfo)
	if err != nil {
		return nil, err
	}

	// change the direct parent's "Children" field
	parentName := pt.Dir(name)
	parentInfo, err := getItemInfoFromMongo(parentName)
	if err != nil {
		return nil, err
	}
	newChildInfo := &pb.ItemBasicInfo{
		Name: name,
		Type: createInfo.Type,
		Size: createInfo.Size,
	}
	parentInfo.Children = append(parentInfo.Children, newChildInfo)
	_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": parentName}, bson.M{"$set": bson.M{"children": parentInfo.Children}})
	if err != nil {
		return nil, err
	}

	return &itemInfo, nil
}

func createItemInDisk(itemInfo *pb.ItemInfo) error {
	if itemInfo.Type == pb.ItemType_File {
		file, err := os.OpenFile(itemInfo.Path, os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
		err = changeSizeAll(itemInfo.Name, itemInfo.Size)
		if err != nil {
			return err
		}
		return nil
	}
	err := os.Mkdir(itemInfo.Path, 0777)
	return err
}

func changeSizeAll(path string, size int64) error {
	// change the size of parents
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			targetPath := path[:i]
			targetInfo, err := getItemInfoFromMongo(targetPath)
			if err != nil {
				return err
			}
			newSize := targetInfo.Size + size
			_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": targetPath}, bson.M{"$set": bson.M{"size": newSize}})
			if err != nil {
				return err
			}
		}
	}

	// change the size of parents' "Children" field
	target := pt.Dir(path)
	for {
		parent := pt.Dir(target)
		if parent == "." {
			break
		}
		parentInfo, err := getItemInfoFromMongo(parent)
		if err != nil {
			return err
		}
		for i, child := range parentInfo.Children {
			if child.Name == target {
				parentInfo.Children[i].Size += size
				break
			}
		}
		_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": parent}, bson.M{"$set": bson.M{"children": parentInfo.Children}})
		if err != nil {
			return err
		}
		target = parent
	}

	return nil
}

func deleteItem(name string) error {
	info, err := getItemInfoFromMongo(name)
	if err != nil {
		return err
	}

	// delete files and directories in disk
	path := info.Path
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	// change parents' size
	size := info.Size
	err = changeSizeAll(name, -size)

	// delete target and all of its children
	err = deleteAllInMongo(name)
	if err != nil {
		return err
	}

	// change the direct parent's "Children" field
	parentName := pt.Dir(name)
	parentInfo, err := getItemInfoFromMongo(parentName)
	if err != nil {
		return err
	}
	for index := 0; index < len(parentInfo.Children); index++ {
		if parentInfo.Children[index].Name == name {
			parentInfo.Children = append(parentInfo.Children[:index], parentInfo.Children[index+1:]...)
			break
		}
	}
	_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": parentName}, bson.M{"$set": bson.M{"children": parentInfo.Children}})
	return err
}

func deleteAllInMongo(name string) error {
	info, err := getItemInfoFromMongo(name)
	if err != nil {
		return err
	}
	for _, child := range info.Children {
		err = deleteAllInMongo(child.Name)
		if err != nil {
			return err
		}
	}
	_, err = db.Collection(itemsCollectionName).DeleteOne(context.Background(), bson.M{"name": name})
	return err
}

func renameItem(name string, newItemName string) error {
	if len(newItemName) < FilenameMinLength {
		return errors.New("file name is too short")
	}
	if len(newItemName) > FilenameMaxLength {
		return errors.New("file name is too long")
	}
	newName := pt.Dir(name) + "/" + newItemName
	if len(newName) > FilepathMaxLength {
		return errors.New("file path name is too long")
	}
	info, err := getItemInfoFromMongo(name)
	if err != nil {
		return err
	}
	newPath := pt.Dir(info.Path) + "/" + base64Encode(newItemName)
	_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": name}, bson.M{"$set": bson.M{"name": newName, "path": newPath}})
	if err != nil {
		panic(err)
	}

	// change the direct parent's "Children" field
	parentName := pt.Dir(name)
	parentInfo, err := getItemInfoFromMongo(parentName)
	if err != nil {
		return err
	}
	for index := 0; index < len(parentInfo.Children); index++ {
		if parentInfo.Children[index].Name == name {
			parentInfo.Children[index].Name = newName
			break
		}
	}
	_, err = db.Collection(itemsCollectionName).UpdateOne(context.Background(), bson.M{"name": parentName}, bson.M{"$set": bson.M{"children": parentInfo.Children}})
	return err
}
