CLI Todo List

A simple and efficient Command Line Interface (CLI) Todo List application built using Go.
This project helps users manage their daily tasks directly from the terminal with ease and speed.

📌 Features

➕ Add new tasks

📋 View all tasks

✅ Mark tasks as completed

❌ Delete tasks

💾 Persistent storage (tasks saved in a file)

⚡ Fast and lightweight CLI tool

🛠️ Tech Stack

Language: Go (Golang)

CLI Framework: Cobra

Data Storage: Local file (JSON / text-based)

📂 Project Structure
CLI_TODO_LIST/
│
├── cmd/              # Cobra commands
│   ├── add.go
│   ├── list.go
│   ├── delete.go
│   └── root.go
│
├── data/             # Task storage
│   └── todo.json
│
├── main.go           # Application entry point
├── go.mod
└── README.md


(Structure may vary slightly depending on implementation)

🚀 Installation & Usage
1️⃣ Clone the Repository
git clone https://github.com/S-Unknown047/CLI_TODO_LIST.git
cd CLI_TODO_LIST

2️⃣ Install Dependencies
go mod tidy

3️⃣ Run the Application
go run main.go


Or build it:

go build
./CLI_TODO_LIST

📖 Commands
Command	Description
add "task name"	Add a new task
list	Show all tasks
done <id>	Mark a task as completed
delete <id>	Delete a task

Example:

todo add "Learn Go"
todo list
todo done 1
todo delete 1

📁 Data Storage

Tasks are stored locally in a file so that data is not lost between executions.

Example task format:

{
  "id": 1,
  "task": "Learn Go",
  "completed": false
}

🎯 Future Improvements

🔔 Due dates & priorities

🔍 Search tasks

🗂️ Categories / tags

☁️ Cloud or database storage

🧪 Unit testing

🤝 Contributing

Contributions are welcome!

Fork the repository

Create a new branch

Make your changes

Submit a pull request

📜 License

This project is open-source and available under the MIT License.

👤 Author

S-Unknown047
GitHub: https://github.com/S-Unknown047

If you want, I can:

🔧 Customize this README exactly to your code

✨ Make it shorter (for hackathon)

📦 Add badges (Go version, license, build)

https://roadmap.sh/projects/task-tracker
https://roadmap.sh/projects/task-tracker
