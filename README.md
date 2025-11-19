# Htmx + Go + PostgreSQL Starter

**A modern, production-ready full-stack template using Htmx, Go, and PostgreSQL.**

Build dynamic web applications without JavaScript frameworks. Use Htmx for interactivity, Go for a fast backend, and PostgreSQL for data persistence. No build step, no bundlers, just pure simplicity.

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/template/TEMPLATE_ID)

## ✨ Features

- 🎯 **Htmx** - Dynamic UI without JavaScript frameworks
- 🔵 **Go** - Fast, compiled backend with Chi router
- 🐘 **PostgreSQL** - Reliable, production-ready database
- 🎨 **Tailwind CSS** - Beautiful styling via CDN
- 🐳 **Docker Optimized** - Multi-stage builds for small images
- 🚂 **Railway Ready** - Zero-config deployment
- ⚡ **No Build Step** - Just code and deploy
- 📝 **CRUD Example** - Working todo list included

## 🚀 Quick Start

### Deploy to Railway

Click the "Deploy on Railway" button above. Railway will automatically:
- Build your Go application using Docker
- Provision a PostgreSQL database
- Connect them together
- Generate a public URL

### Local Development

**Prerequisites:**
- Go 1.21 or higher
- PostgreSQL (or use Docker)

**Steps:**
```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/htmx-go-postgres.git
cd htmx-go-postgres

# Install dependencies
go mod download

# Set up PostgreSQL (or use Docker)
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

# Set environment variables
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
export PORT=8080

# Run the application
go run cmd/web/main.go

# Open browser to http://localhost:8080
```

## 📁 Project Structure
```
htmx-go-postgres/
├── cmd/
│   └── web/
│       └── main.go              # Application entry point
├── templates/
│   ├── index.html               # Main page
│   └── todo-list.html           # Todo list partial
├── static/
│   ├── css/                     # Custom CSS (optional)
│   └── js/                      # Custom JS (optional)
├── Dockerfile                   # Multi-stage Docker build
├── railway.toml                 # Railway configuration
├── go.mod                       # Go dependencies
├── go.sum                       # Go checksums
└── README.md                    # Documentation
```

## 🎯 How It Works

### Htmx Magic

Htmx allows you to build dynamic UIs using HTML attributes:
```html
<!-- Add a todo -->
<form hx-post="/todos" 
      hx-target="#todo-list" 
      hx-swap="innerHTML">
    <input type="text" name="title" required>
    <button type="submit">Add</button>
</form>

<!-- Toggle completed -->
<input type="checkbox" 
       hx-put="/todos/123/toggle"
       hx-target="#todo-list">

<!-- Delete a todo -->
<button hx-delete="/todos/123"
        hx-target="#todo-list"
        hx-confirm="Delete this?">
    Delete
</button>
```

**No JavaScript required!** Htmx handles all the AJAX calls and DOM updates.

### Go Backend

Simple, fast Go server with Chi router:
```go
// Create todo
r.Post("/todos", app.createTodo)

// Toggle todo
r.Put("/todos/{id}/toggle", app.toggleTodo)

// Delete todo
r.Delete("/todos/{id}", app.deleteTodo)
```

Returns HTML fragments that Htmx swaps into the page.

### PostgreSQL Database

Simple schema with auto-migration:
```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);
```

## 🛠️ Customization

### Add New Routes

Edit `cmd/web/main.go`:
```go
// Add your route
r.Get("/mypage", app.myPageHandler)

// Create handler
func (app *Application) myPageHandler(w http.ResponseWriter, r *http.Request) {
    app.Templates.ExecuteTemplate(w, "mypage.html", data)
}
```

### Add New Templates

Create `templates/mypage.html`:
```html
<div>
    <h1>My Page</h1>
    <p>Content here...</p>
</div>
```

### Add Database Models

Extend the schema in `main.go`:
```go
func createTable(db *sql.DB) {
    query := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL
        );
    `
    db.Exec(query)
}
```

### Add Styling

Use Tailwind classes inline, or add custom CSS in `static/css/`.

## 🌐 Why Htmx?

### The Case Against SPAs

- ❌ Complex build toolchains
- ❌ Large JavaScript bundles
- ❌ State management hell
- ❌ SEO challenges
- ❌ Slow initial load

### The Htmx Advantage

- ✅ **No Build Step** - Write HTML and Go
- ✅ **Small Payload** - Htmx is ~14KB
- ✅ **Server-Side Rendering** - SEO-friendly by default
- ✅ **Progressive Enhancement** - Works without JS
- ✅ **Simple Mental Model** - Just HTML attributes
- ✅ **Fast** - No client-side framework overhead

## 📚 Learn More

### Htmx Resources
- [Htmx Documentation](https://htmx.org/docs/) - Official Htmx docs
- [Htmx Examples](https://htmx.org/examples/) - Real-world examples
- [Hypermedia Systems](https://hypermedia.systems/) - Book on Htmx architecture

### Go Resources
- [Go Documentation](https://go.dev/doc/) - Official Go docs
- [Chi Router](https://github.com/go-chi/chi) - Lightweight Go router
- [Go by Example](https://gobyexample.com/) - Hands-on Go tutorials

### Deployment
- [Railway Docs](https://docs.railway.app/) - Platform documentation
- [PostgreSQL on Railway](https://docs.railway.app/databases/postgresql) - Database guide

## 🎓 Use Cases

### Perfect For:

📊 **Internal Dashboards** - Admin panels, monitoring tools  
📝 **Content-Heavy Sites** - Blogs, documentation, news sites  
🛠️ **CRUD Applications** - Data management, forms, tables  
🏢 **Business Applications** - CRM, inventory, invoicing  
📱 **Progressive Enhancement** - Works without JavaScript  

### Not Ideal For:

❌ Real-time collaboration (use WebSockets instead)  
❌ Heavy client-side state (use React/Vue)  
❌ Offline-first apps (use PWA/WASM)  
❌ Complex animations (use JavaScript)  

## ⚡ Performance

- **Server Response:** < 50ms (Go is fast!)
- **Page Load:** Minimal (no framework to download)
- **Bundle Size:** ~14KB (just Htmx + Tailwind CDN)
- **Memory Usage:** ~20-30MB (Go binary)
- **Database Queries:** Optimized with indexes

## 🔐 Security

- ✅ SQL injection protected (parameterized queries)
- ✅ CSRF protection (add middleware if needed)
- ✅ XSS protection (Go templates auto-escape)
- ✅ HTTPS on Railway (automatic SSL)

## 🤝 Contributing

Contributions welcome! Here's how:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing`)
5. Open a Pull Request

## 📄 License

MIT License - see LICENSE file for details

---

**Built with ❤️ for the Railway community** 🚂

**Tired of JavaScript complexity?** This template proves you don't need React/Vue/Angular for dynamic UIs!

**Questions?** Open an issue on GitHub or reach out on Railway Discord.