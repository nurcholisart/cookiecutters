# {{ cookiecutter.project_name }}
{{ cookiecutter.project_description }}

### Requirements
- Go 1.17
- Qiscus Multichannel Account

### Local Setup
**Clone Repository**

```bash
git clone `project repository`
cd `directory`
```

**Install Packages**
```bash
go mod tidy
```

**Setup Environment Variables**
> Please check the `.env.sample` for the description of each variables
```bash
cp .env.sample .env
```

**Run Test**

```bash
go test ./...
```

**Run Service**
```bash
go run cmd/main.go
```
