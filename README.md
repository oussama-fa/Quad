# Quad: Pattern Generator in Go

Welcome to **Quad**, a beginner-friendly project showcasing various implementations of rectangular pattern generation in **Go (Golang)**. This project was developed during a one-week study at **Zone01 Oujda**, focusing on **DevOps** concepts and practical programming.

---

## 📚 Project Overview

The `Quad` project contains several functions (`QuadA`, `QuadB`, `QuadC`, `QuadD`, `QuadE`) that generate distinct rectangular patterns using different characters. Each pattern highlights an understanding of nested loops, conditional logic, and character-based design in Go.

### **Key Features**:
- Generates various rectangular patterns using characters like `o`, `A`, `B`, `C`, `*`, and more.
- Demonstrates beginner-friendly Go programming concepts.
- Utilizes the external library [z01](https://pkg.go.dev/github.com/01-edu/z01) for character printing.

---

## 🔧 How It Works

Each function (`QuadA`, `QuadB`, etc.) accepts two parameters:
- `x`: Width of the rectangle.
- `y`: Height of the rectangle.

Based on the input dimensions, the program prints rectangular patterns with specific border and fill characters.  

### Example Outputs:
- For `QuadA(5, 3)`:
	```bash
	o---o
	|   |
	o---o
	```

- For `QuadC(5, 3)`:
	```bash
	ABBBA
	B   B
	CBBBC
	```
---

## 🚀 Getting Started

### Prerequisites
- Install [Go](https://go.dev/) (version 1.16 or higher).
- Clone this repository:  
  ```bash
  git clone https://github.com/oussama-fa/Quad.git
  cd Quad

### Ensure the z01 library is installed. You can get it via:
- ```bash
	go get github.com/01-edu/z01

### Running the Program:
- ```bash
	touch main.go

- Edit main.go and include the following code:
	```go
	package main

	import "your-repo-path/piscine"

	func main() {
		piscine.QuadA(5, 3)
		piscine.QuadC(7, 4)
	}
	
- RUN
	go run main.go

## 🛠️ Tech Stack
* Language: Go (Golang)
* DevOps Tools: Concepts learned from Zone01 Oujda
* External Libraries: z01


This version is a **single continuous section** from start to finish. You can now directly copy and paste it into your `README.md`. Let me know if you need further tweaks!
