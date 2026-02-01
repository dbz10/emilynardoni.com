# Agent Guidelines for emilynardoni.com

## Overview
- Go web app using `net/http` with server code in `cmd/web`.
- Server renders HTML templates from `static/html` and serves assets from `static/`.
- No database or API layer; content is mostly static and lives in templates and static files.
- Deployment uses Fly.io (see `README.md` for `fly launch` / `fly deploy`).

## Repository Structure
- `cmd/web`: HTTP server entrypoint, routes, handlers, template loading/cache.
- `cmd/importer`: standalone utility (not used by the web server runtime).
- `pkg/importer`: shared code for the importer.
- `static/html`: Go HTML templates (pages, partials, layouts).
- `static/css`: site styles.
- `static/img` and `static/content`: images and PDFs.
- `static/embed.go`: embeds static assets for builds.

## Web App Conventions
- Routing lives in `cmd/web/routes.go` using `http.NewServeMux()`.
- Each route has a handler method on `application` in `cmd/web/handlers.go`.
- Templates:
  - Page templates are `*.page.tmpl` and include `base` layout.
  - Partials are `*.partial.tmpl` (e.g., navbar/footer).
  - Layouts are `*.layout.tmpl` (e.g., `base.layout.tmpl`).
- Template cache is built once at startup in `cmd/web/main.go` using `newTemplateCache`.
- The `render` helper in `cmd/web/handlers.go` always provides `GenericData` with `Now` for template use.

## Adding a New Page
1) Create `static/html/<name>.page.tmpl` that includes `{{ template "base" . }}`.
2) Add a handler in `cmd/web/handlers.go` following existing pattern.
3) Register the route in `cmd/web/routes.go`.
4) Add the nav link in `static/html/navbar.partial.tmpl` if needed.

## Styling Conventions
- Global styles in `static/css/style.css`.
- Fonts: Google Fonts (Montserrat, Belleza) and local `roca-regular.ttf`.
- Typography and headings use CSS classes like `.subheading` and `.photo-caption`.

## Asset Conventions
- Use `/static/...` paths in templates for images, CSS, and PDFs.
- PDFs are stored in `static/content`.
- When updating a static asset without renaming the file, bump its cache-busting query parameter (e.g., `?v=YYYYMMDD`) in templates so browsers fetch the new version.

## Development Notes
- The server listens on port `:8080` and uses the template cache by default.
- There is no test suite in this repo; changes should be verified manually in the browser.
