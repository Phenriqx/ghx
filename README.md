# Github CLI Tool
A command-line interface (CLI) tool built with Golang to interact with GitHub repositories efficiently. This tool leverages Cobra for command handling and Resty for making API requests.


## 🚀 Features
<ul>
  
- 🔧 **Initialize Repositories**  
  Automatically create a GitHub repository, set remote origin, and optionally add README, .gitignore, and push.
  
- 📜 **List Repositories**  
  View all your GitHub repositories in a clean list.
  
- 📦 **Create New Repositories**  
  Quickly create public or private repos with optional descriptions.

- 🧹 **Delete Repositories**  
  Permanently delete repositories with confirmation.

- 🔍 **Repository Info**  
  Fetch metadata and details for a given repository.

- 🧲 **Clone Repositories**  
  Clone repos using HTTPS or SSH.

- 🔁 **Manage Pull Requests**  
  Create, list, merge, and close PRs with ease.

- 🎮 **Interactive Mode (coming soon)**  
  Use `ghx init` with no arguments to launch an interactive setup wizard.
</ul>



## 🛠 Installation 
Make sure you have Go installed on your system. Then, clone the repository and build the CLI tool:

```
# Clone the repository
git clone https://github.com/yourusername/ghx.git
cd ghx

# Build the binary
go build -o ghx .

# Move to a directory in your PATH (optional)
sudo mv ghx /usr/local/bin/ (linux/macOS)
```
After that, you need to make your compiled binary executable by running chmod:

```
chmod +x /usr/local/bin/ghx
```

## 🔑 Configuration
This tool requires a GitHub personal access token for authentication. You can set up your token as an environment variable like this:
```
export GITHUB_TOKEN="your_personal_access_token"
```

Or you can create a .env file in the same path as your compiled code (recommended):

```
echo "GITHUB_TOKEN=your_actual_token_here" > .env
```

## 🏃 Usage
Run the CLI tool with the available commands:

```
ghx --help
```
### 📖 Example Commands
```
ghx list
```
```
ghx create <repo-name> --private (optional) --desc "" (optional)
```
```
ghx delete <repo-name>
```


## 📦 Dependencies 
<p>🐍 <a href="https://github.com/spf13/cobra" >Cobra</a> - CLI framework</p>
<p>⚡ <a href="https://github.com/go-resty/resty" style="text-decoration:none;">Resty</a> - HTTP client for API requests</p>

