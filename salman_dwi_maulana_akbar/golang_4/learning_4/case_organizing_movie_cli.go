package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	Title    string
	Genre    string
	Director string
	Year     int
	Rating   float64
}

var movieCollection = make(map[string]Movie)

func addMovie() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter movie title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter movie genre: ")
	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	fmt.Print("Enter movie director: ")
	director, _ := reader.ReadString('\n')
	director = strings.TrimSpace(director)

	// Validate year input
	var year int
	for {
		fmt.Print("Enter release year: ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)
		var err error
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer for the year.")
		} else {
			break
		}
	}

	// Validate rating input
	var rating float64
	for {
		fmt.Print("Enter rating: ")
		ratingStr, _ := reader.ReadString('\n')
		ratingStr = strings.TrimSpace(ratingStr)
		var err error
		rating, err = strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid float number for the rating.")
		} else {
			break
		}
	}

	movie := Movie{
		Title:    title,
		Genre:    genre,
		Director: director,
		Year:     year,
		Rating:   rating,
	}

	movieCollection[title] = movie
	fmt.Println("Movie added successfully!")
}

func deleteMovie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter movie title to delete: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if _, exists := movieCollection[title]; exists {
		delete(movieCollection, title)
		fmt.Println("Movie deleted successfully!")
	} else {
		fmt.Println("Movie not found!")
	}
}

func updateMovie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter movie title to update: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if movie, exists := movieCollection[title]; exists {
		fmt.Print("Enter new genre (leave blank to keep current): ")
		genre, _ := reader.ReadString('\n')
		genre = strings.TrimSpace(genre)
		if genre != "" {
			movie.Genre = genre
		}

		fmt.Print("Enter new director (leave blank to keep current): ")
		director, _ := reader.ReadString('\n')
		director = strings.TrimSpace(director)
		if director != "" {
			movie.Director = director
		}

		fmt.Print("Enter new release year (leave blank to keep current): ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)
		if yearStr != "" {
			year, err := strconv.Atoi(yearStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer for the year.")
			} else {
				movie.Year = year
			}
		}

		fmt.Print("Enter new rating (leave blank to keep current): ")
		ratingStr, _ := reader.ReadString('\n')
		ratingStr = strings.TrimSpace(ratingStr)
		if ratingStr != "" {
			rating, err := strconv.ParseFloat(ratingStr, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid float number for the rating.")
			} else {
				movie.Rating = rating
			}
		}

		movieCollection[title] = movie
		fmt.Println("Movie updated successfully!")
	} else {
		fmt.Println("Movie not found!")
	}
}

func searchMovie() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter movie title to search: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if movie, exists := movieCollection[title]; exists {
		fmt.Printf("Title: %s\nGenre: %s\nDirector: %s\nYear: %d\nRating: %.1f\n",
			movie.Title, movie.Genre, movie.Director, movie.Year, movie.Rating)
	} else {
		fmt.Println("Movie not found!")
	}
}

func listMovies() {
	if len(movieCollection) == 0 {
		fmt.Println("No movies in the collection.")
		return
	}

	index := 1
	for _, movie := range movieCollection {
		fmt.Printf("%d: Title: %s\tGenre: %s\tDirector: %s\tYear: %d\tRating: %.1f\n",
			index, movie.Title, movie.Genre, movie.Director, movie.Year, movie.Rating)
		index++
	}
}

func saveToFile() {
	file, err := os.Create("movies.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(movieCollection)
	if err != nil {
		fmt.Println("Error saving data:", err)
	} else {
		fmt.Println("Movies saved to file successfully!")
	}
}

func loadFromFile() {
	file, err := os.Open("movies.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&movieCollection)
	if err != nil {
		fmt.Println("Error loading data:", err)
	} else {
		fmt.Println("Movies loaded from file successfully!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMovie Manager")
		fmt.Println("1. Add Movie")
		fmt.Println("2. Delete Movie")
		fmt.Println("3. Update Movie")
		fmt.Println("4. Search Movie")
		fmt.Println("5. List Movies")
		fmt.Println("6. Save to File")
		fmt.Println("7. Load from File")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(choiceStr))

		switch choice {
		case 1:
			addMovie()
		case 2:
			deleteMovie()
		case 3:
			updateMovie()
		case 4:
			searchMovie()
		case 5:
			listMovies()
		case 6:
			saveToFile()
		case 7:
			loadFromFile()
		case 8:
			fmt.Println("Exiting Movie Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
