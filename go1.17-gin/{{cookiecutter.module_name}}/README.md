# {{ cookiecutter.project_name }}
{{ cookiecutter.project_description }}

### Requirements
- Go {{ cookiecutter.go_version }}
- Qiscus Multichannel Account

### Local Setup
**Clone Repository**

```bash
git clone {{ project repository }}
cd {{ project repository }}
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
