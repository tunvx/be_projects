package util

import (
	"math/rand"
	"strings"
	"time"
)

var sampleAuthors = []string{
	"Haruki Murakami", "George Orwell", "J.K. Rowling", "Yuval Noah Harari",
	"Nguyễn Nhật Ánh", "Ernest Hemingway", "Jane Austen", "Stephen King",
	"Agatha Christie", "Mark Twain", "F. Scott Fitzgerald", "J.R.R. Tolkien",
	"Isaac Asimov", "Ray Bradbury", "Margaret Atwood", "Gabriel García Márquez",
	"Leo Tolstoy", "Virginia Woolf", "Herman Melville", "Charles Dickens",
	"Fyodor Dostoevsky", "Oscar Wilde", "Emily Dickinson", "Sylvia Plath",
	"John Steinbeck", "Kurt Vonnegut", "Toni Morrison", "Chimamanda Ngozi Adichie",
	"Neil Gaiman", "C.S. Lewis", "Philip", "Dickens", "Arthur Conan Doyle",
}

var samplePublishers = []string{
	"Penguin Books", "HarperCollins", "Nhà Xuất Bản Trẻ", "Simon & Schuster",
	"O'Reilly Media", "Random House", "Macmillan Publishers", "Hachette Livre",
	"Oxford University Press", "Cambridge University Press",
	"Springer Nature", "Wiley", "Pearson Education", "Scholastic Inc.",
	"University of Chicago Press", "MIT Press", "Princeton University Press",
	"Cambridge University Press", "John Wiley & Sons", "Taylor & Francis",
	"Elsevier", "SAGE Publications", "Routledge", "CRC Press",
	"Bloomsbury Publishing", "Farrar, Straus and Giroux", "Knopf Doubleday",
	"Little, Brown and Company", "St. Martin's Press", "Doubleday",
	"Harper Perennial", "Vintage Books", "Bantam Books", "Ballantine Books",
	"Tor Books", "Orbit Books", "Ace Books", "DAW Books",
	"Del Rey", "Gollancz", "Head of Zeus", "Quercus Publishing",
	"Pan Macmillan", "Hodder & Stoughton", "Bloomsbury Children's Books",
}

var sampleCategories = []string{
	"Fiction", "Science", "Fantasy", "Biography", "Philosophy", "Children", "Mystery",
}

var sampleTags = []string{
	"bestseller", "classic", "2024", "2025", "new release", "award-winning",
}

var sampleEditions = []string{
	"1st edition", "2nd edition", "3th edition", "4th edition", "5th edition",
	"6th edition", "7th edition", "8th edition", "9th edition", "10th edition",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func RandomAuthor() string {
	return sampleAuthors[rand.Intn(len(sampleAuthors))]
}

func RandomPublisher() string {
	return samplePublishers[rand.Intn(len(samplePublishers))]
}

func RandomEdition() string {
	return sampleEditions[rand.Intn(len(sampleEditions))]
}

func RandomReleaseDate() string {
	year := rand.Intn(40) + 1985 // 1985 - 2024
	month := rand.Intn(12) + 1
	day := rand.Intn(28) + 1
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
}

func RandomDescription() string {
	descriptions := []string{
		"A gripping tale of adventure and discovery.",
		"An insightful look into the human condition.",
		"A thrilling mystery that will keep you guessing.",
		"A heartwarming story about love and resilience.",
		"A thought-provoking exploration of society and culture.",
		"A captivating fantasy world filled with magic and wonder.",
		"A historical epic that brings the past to life.",
		"A chilling horror story that will haunt your dreams.",
		"A humorous take on everyday life and its absurdities.",
		"A poetic narrative that delves into the depths of emotion.",
		"A science fiction journey through time and space.",
		"A philosophical treatise on existence and meaning.",
		"A collection of short stories that explore various themes.",
		"A memoir that shares personal experiences and insights.",
		"A biography that chronicles the life of a remarkable individual.",
		"A self-help guide to personal growth and development.",
		"A travelogue that takes you to distant lands and cultures.",
		"A cookbook filled with delicious recipes and culinary tips.",
		"A graphic novel that combines art and storytelling.",
		"A children's book that sparks imagination and creativity.",
		"A classic novel that has stood the test of time.",
		"A modern retelling of a timeless story.",
		"A dystopian narrative that warns of future possibilities.",
		"A romance that explores the complexities of love.",
		"A satire that critiques societal norms and expectations.",
		"A collection of essays that offer unique perspectives.",
		"A motivational book that inspires action and change.",
		"A guide to mastering a specific skill or craft.",
		"A deep dive into a specific topic or field of study.",
		"A collection of poems that evoke strong emotions.",
		"A narrative that intertwines multiple storylines.",
		"A suspenseful thriller that keeps you on the edge of your seat.",
		"A heartwarming children's story that teaches valuable lessons.",
		"A gripping historical fiction that immerses you in the past.",
	}
	return descriptions[rand.Intn(len(descriptions))]
}

func RandomPageCount() int {
	return rand.Intn(800) + 100 // 100 - 899 pages
}

func RandomContent() string {
	return RandomString(10000) // Placeholder for book content
}

func RandomCategories() []string {
	n := rand.Intn(2) + 1 // 1-2 categories
	m := make(map[string]struct{})
	for len(m) < n {
		m[sampleCategories[rand.Intn(len(sampleCategories))]] = struct{}{}
	}
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}

func RandomTags() []string {
	n := rand.Intn(3) + 1 // 1-3 tags
	m := make(map[string]struct{})
	for len(m) < n {
		m[sampleTags[rand.Intn(len(sampleTags))]] = struct{}{}
	}
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}

func RandomRating() float32 {
	return float32(rand.Intn(21)+30) / 10 // 3.0 - 5.0
}

func RandomReviewCount() int {
	return rand.Intn(5000) // 0 - 4999 reviews
}
