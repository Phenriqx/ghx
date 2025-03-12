# Github CLI Tool
A command-line interface (CLI) tool built with Golang to interact with GitHub repositories efficiently. This tool leverages Cobra for command handling and Resty for making API requests.


## 🚀 Features 

<ul>
  <li>List user repositories</li>
  <li>Fetch repository details</li>
  <li>Create new repositories</li>
  <li>Clone repositories</li>
  <li>Delete repositories</li>
</ul>


## 🛠 Installation 
Make sure you have Go installed on your system. Then, clone the repository and build the CLI tool:

```
# Clone the repository
git clone https://github.com/yourusername/github-cli.git
cd github-cli

# Build the binary
go build -o github-cli

# Move to a directory in your PATH (optional)
sudo mv github-cli /usr/local/bin/ (linux/macOS)
```

## 🏃 Usage
Run the CLI tool with the available commands:

```
github-cli --help
```
### 📖 Example Commands
```
github-cli list
```
```
github-cli create <repo-name> --private (optional) --desc "" (optional)
```
```
github-cli delete <repo-name>
```

## 🔑 Configuration
This tool requires a GitHub personal access token for authentication. Set up your token as an environment variable:
```
export GITHUB_TOKEN="your_personal_access_token"
```

## 📦 Dependencies 
<p>🐍 <a href="https://github.com/spf13/cobra" >Cobra</a> - CLI framework</p>
<p>⚡ <a href="https://github.com/go-resty/resty" style="text-decoration:none;">Resty</a> - HTTP client for API requests</p>

