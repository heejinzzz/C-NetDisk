package operation

import "fmt"

type ProcessBar struct {
	Header        string
	TotalSize     int64
	Width         int
	CurrentString string
}

func NewProcessBar(header string, totalSize int64) *ProcessBar {
	return &ProcessBar{
		Header:    header,
		TotalSize: totalSize,
		Width:     DefaultProcessBarWidth,
	}
}

func (bar *ProcessBar) SetWidth(width int) {
	bar.Width = width
}

func (bar *ProcessBar) ShowProcess(currentSize int64) {
	fmt.Printf("\r")

	p := float32(currentSize) / float32(bar.TotalSize)
	count := p * float32(bar.Width)
	tmp := make([]byte, int(count))
	for i := 0; i < int(count); i++ {
		tmp[i] = '='
	}

	currentUnitIndex := 0
	size := float64(currentSize)
	for size >= 1024 && currentUnitIndex < len(SizeUnits)-1 {
		size /= 1024.0
		currentUnitIndex++
	}

	totalUnitIndex := 0
	totalSize := float32(bar.TotalSize)
	for totalSize >= 1024 && totalUnitIndex < len(SizeUnits)-1 {
		totalSize /= 1024.0
		totalUnitIndex++
	}
	processSize := fmt.Sprintf("%.2f", size) + SizeUnits[currentUnitIndex] + "/" + fmt.Sprintf("%.2f", totalSize) + SizeUnits[totalUnitIndex]

	bar.CurrentString = fmt.Sprintf("%-45s | %-16s| %s %d%%", bar.Header, processSize, string(tmp), int(p*100))
	fmt.Printf("%s", bar.CurrentString)
	if bar.Width == int(count) {
		fmt.Println("")
	}
}

func NewProcessBarHeader(command string, target string) string {
	if len(target) > ItemNameMaxInvisibleLength {
		target = target[:ItemNameMaxInvisibleLength] + "..."
	}
	header := command + " [" + target + "]"
	return header
}
