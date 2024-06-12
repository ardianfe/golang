package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fumiama/go-docx"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

type Deck []string

func (d Deck) print() {
	if len(d) <= 0 {
		fmt.Println("Deck belum dibuat, silahkan buat terlebih dahulu")
		return
	}
	for i, card := range d {
		fmt.Println(i+1, ".", card)
	}
}

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename, format string) error {
	switch format {
	case "txt":
		return d.saveToText(filename)
	case "pdf":
		return d.saveToPDF(filename)
	case "excel":
		return d.saveToExcel(filename)
	case "html":
		return d.saveToHTML(filename)
	case "word":
		return d.saveToWord(filename)
	default:
		return fmt.Errorf("format tidak dikenal: %s", format)
	}
}

func (d Deck) saveToText(filename string) error {
	directory := "file_save_deck/txt"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	return os.WriteFile("file_save_deck/txt/"+filename+".txt", []byte(d.toString()), 0666)
}

func (d Deck) saveToPDF(filename string) error {
	directory := "file_save_deck/pdf"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	for i, card := range d {
		pdf.Cell(40, 10, fmt.Sprintf("%d. %s", i+1, card))
		pdf.Ln(12)
	}
	return pdf.OutputFileAndClose("file_save_deck/pdf/" + filename + ".pdf")
}

func (d Deck) saveToExcel(filename string) error {
	directory := "file_save_deck/excel"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	f := xlsx.NewFile()
	sheet, err := f.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	for i, card := range d {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%d. %s", i+1, card)
	}

	return f.Save("file_save_deck/excel/" + filename + ".xlsx")
}

func (d Deck) saveToHTML(filename string) error {
	directory := "file_save_deck/html"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	htmlContent := "<html><body><ul>"
	for _, card := range d {
		htmlContent += "<li>" + card + "</li>"
	}
	htmlContent += "</ul></body></html>"
	return os.WriteFile("file_save_deck/html/"+filename+".html", []byte(htmlContent), 0666)
}

func (d Deck) saveToWord(filename string) error {
	directory := "file_save_deck/word"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	var w *docx.Docx
	w = docx.New().WithDefaultTheme().WithA4Page()

	for _, card := range d {
		w.AddParagraph().AddText(card)
	}

	f, err := os.Create("file_save_deck/word/" + filename + ".docx")
	if err != nil {
		return err
	}

	_, err = w.WriteTo(f)
	// Save document to file
	return f.Close()
}

func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ";")
	return Deck(s)
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Deck) addItem(item string) Deck {
	for _, listItem := range d {
		if listItem == item {
			fmt.Println("Item", item, "sudah ada di deck.")
			return d
		}
	}
	d = append(d, item)
	fmt.Println("Item", item, "telah ditambahkan ke deck.")
	return d
}

func (d Deck) deleteItem(item string) Deck {
	for index, listItem := range d {
		if listItem == item {
			d = append(d[:index], d[index+1:]...)
			fmt.Println("Item", item, "telah dihapus dari deck.")
			return d
		}
	}
	fmt.Println("Item", item, "tidak ditemukan di deck.")
	return d
}

func (d Deck) updateItem(oldItem, newItem string) Deck {
	for index, listItem := range d {
		if listItem == oldItem {
			d[index] = newItem
			fmt.Println("Item", oldItem, "telah diubah menjadi", newItem)
			return d
		}
	}
	fmt.Println("Item", oldItem, "tidak ditemukan di deck.")
	return d
}

func (d Deck) searchItem(item string) {
	found := false
	for index, listItem := range d {
		if listItem == item {
			fmt.Printf("%s ditemukan pada indeks %d\n", item, index+1)
			found = true
		}
	}
	if !found {
		fmt.Printf("%s tidak ditemukan di deck\n", item)
	}
}

func (d Deck) sortList() {
	sort.Strings(d)
	fmt.Println("Deck telah diurutkan:")
	d.print()
}


func main() {
	var deck Deck
	reader := bufio.NewReader(os.Stdin)
	var choice int
	var item, newItem, oldItem, filename, format string

	for {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Buat deck baru")
		fmt.Println("2. Tambahkan item ke deck")
		fmt.Println("3. Hapus item dari deck")
		fmt.Println("4. Ubah item di deck")
		fmt.Println("5. Lihat deck")
		fmt.Println("6. Cari item di deck")
		fmt.Println("7. Urutkan deck")
		fmt.Println("8. Shuffle deck")
		fmt.Println("9. Simpan deck ke file")
		fmt.Println("10. Baca deck dari file")
		fmt.Println("11. Keluar")
		fmt.Print("Pilihan Anda: ")
		input, _ := reader.ReadString('\n')
		choice, _ = strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			deck = newDeck()
			fmt.Println("Deck baru telah dibuat.")
		case 2:
			fmt.Print("Masukkan item baru: ")
			item, _ = reader.ReadString('\n')
			item = strings.TrimSpace(item)
			deck = deck.addItem(item)
		case 3:
			fmt.Print("Masukkan nama item yang ingin dihapus: ")
			item, _ = reader.ReadString('\n')
			item = strings.TrimSpace(item)
			deck = deck.deleteItem(item)
		case 4:
			fmt.Print("Masukkan nama item yang ingin diubah: ")
			oldItem, _ = reader.ReadString('\n')
			oldItem = strings.TrimSpace(oldItem)
			fmt.Print("Masukkan item baru: ")
			newItem, _ = reader.ReadString('\n')
			newItem = strings.TrimSpace(newItem)
			deck = deck.updateItem(oldItem, newItem)
		case 5:
			deck.print()
		case 6:
			fmt.Print("Masukkan item yang ingin dicari: ")
			item, _ = reader.ReadString('\n')
			item = strings.TrimSpace(item)
			deck.searchItem(item)
		case 7:
			deck.sortList()
		case 8:
			deck.shuffle()
			fmt.Println("Deck telah di-shuffle.")
		case 9:
			fmt.Print("Masukkan format file (txt, pdf, excel, html, word): ")
			format, _ = reader.ReadString('\n')
			format = strings.TrimSpace(format)
			fmt.Print("Masukkan nama file untuk menyimpan deck: ")
			filename, _ = reader.ReadString('\n')
			filename = strings.TrimSpace(filename)
			err := deck.saveToFile(filename, format)
			if err != nil {
				fmt.Println("Gagal menyimpan deck:", err)
			} else {
				fmt.Println("Deck telah disimpan ke file", filename)
			}
		case 10:
			baseDir := "file_save_deck"
			var allFiles []string
			var groupedFiles [][]string

			items, err := os.ReadDir(baseDir)
			if err != nil {
				fmt.Println("Gagal membaca direktori:", err)
				continue
			}

			// Read files and group by directory
			for _, item := range items {
				if item.IsDir() {
					subItems, err := os.ReadDir(baseDir + "/" + item.Name())
					if err != nil {
						fmt.Println("Gagal membaca subdirektori:", err)
						continue
					}
					var subFiles []string
					subFiles = append(subFiles, item.Name())
					for _, subItem := range subItems {
						subFiles = append(subFiles, subItem.Name())
					}
					groupedFiles = append(groupedFiles, subFiles)
				} else {
					allFiles = append(allFiles, item.Name())
				}
			}

			// Display the grouped files
			fmt.Println("Daftar file yang tersedia:")
			for _, group := range groupedFiles {
				dirName := group[0]
				fmt.Println(dirName + ":")
				for i := 1; i < len(group); i++ {
					fmt.Printf("  %d. %s\n", i, group[i])
				}
			}

			if len(allFiles) > 0 {
				fmt.Println(baseDir + ":")
				for i, file := range allFiles {
					fmt.Printf("  %d. %s\n", i+1, file)
				}
			}

			// Ask user to select a file
			var chosenDir string
			var chosenFile string
			validChoice := false

			for !validChoice {
				fmt.Print("Masukkan nama direktori yang ingin dibaca: ")
				chosenDir, _ = reader.ReadString('\n')
				chosenDir = strings.TrimSpace(chosenDir)

				// Find the selected directory in groupedFiles
				for _, group := range groupedFiles {
					if group[0] == chosenDir {
						fmt.Print("Masukkan nomor file yang ingin dibaca: ")
						input, _ := reader.ReadString('\n')
						input = strings.TrimSpace(input)
						fileNum, err := strconv.Atoi(input)
						if err != nil || fileNum < 1 || fileNum > len(group)-1 {
							fmt.Println("Nomor file tidak valid. Masukkan nomor yang sesuai.")
						} else {
							chosenFile = group[fileNum]
							validChoice = true
						}
						break
					}
				}

				// Check the base directory for files
				if chosenDir == baseDir {
					fmt.Print("Masukkan nomor file yang ingin dibaca: ")
					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(input)
					fileNum, err := strconv.Atoi(input)
					if err != nil || fileNum < 1 || fileNum > len(allFiles) {
						fmt.Println("Nomor file tidak valid. Masukkan nomor yang sesuai.")
					} else {
						chosenFile = allFiles[fileNum-1]
						validChoice = true
					}
				}

				if !validChoice {
					fmt.Println("Direktori tidak valid. Masukkan nama direktori yang sesuai.")
				}
			}

			fullFilePath := baseDir + "/" + chosenDir + "/" + chosenFile
			if chosenDir == baseDir {
				fullFilePath = baseDir + "/" + chosenFile
			}
			deck = newDeckFromFile(fullFilePath)
			fmt.Println("Deck telah dibaca dari file", fullFilePath)
		case 11:
			fmt.Println("Terima kasih telah menggunakan aplikasi pengelola deck!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang sesuai.")
		}
	}
}
