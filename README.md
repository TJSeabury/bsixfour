# bsixfour
Converts the specified image to a web-ready base64 encoded string,
and saves this to the file and location specified.
If no save location is specified, bsixfour assumes you want to save in the same location as the source file.
bsixfour automatically handles the file extension.

---

## Usage
```
bsixfour -in <inPath> [ -out <outPath> -html [true | false] -raw ]
```

## Exaples
### Source
![Cog](/example/cog.png)

### HTML Output
```bash
bsixfour -in ./cog.png
```

### Result
```html
<img
  src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAEsWlUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4KPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNS41LjAiPgogPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIgogICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iCiAgICB4bWxuczpleGlmPSJodHRwOi8vbnMuYWRvYmUuY29tL2V4aWYvMS4wLyIKICAgIHhtbG5zOnBob3Rvc2hvcD0iaHR0cDovL25zLmFkb2JlLmNvbS9waG90b3Nob3AvMS4wLyIKICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIKICAgIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIgogICAgeG1sbnM6c3RFdnQ9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZUV2ZW50IyIKICAgdGlmZjpJbWFnZUxlbmd0aD0iMzIiCiAgIHRpZmY6SW1hZ2VXaWR0aD0iMzIiCiAgIHRpZmY6UmVzb2x1dGlvblVuaXQ9IjIiCiAgIHRpZmY6WFJlc29sdXRpb249IjEzOC4wIgogICB0aWZmOllSZXNvbHV0aW9uPSIxMzguMCIKICAgZXhpZjpQaXhlbFhEaW1lbnNpb249IjMyIgogICBleGlmOlBpeGVsWURpbWVuc2lvbj0iMzIiCiAgIGV4aWY6Q29sb3JTcGFjZT0iMSIKICAgcGhvdG9zaG9wOkNvbG9yTW9kZT0iMyIKICAgcGhvdG9zaG9wOklDQ1Byb2ZpbGU9InNSR0IgSUVDNjE5NjYtMi4xIgogICB4bXA6TW9kaWZ5RGF0ZT0iMjAyMi0w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"
  alt="cog"
/>

```

### Text Output
```bash
bsixfour -in ./cog.png -html false
```

#### Result
```
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAEsWlUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4KPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNS41LjAiPgogPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIgogICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iCiAgICB4bWxuczpleGlmPSJodHRwOi8vbnMuYWRvYmUuY29tL2V4aWYvMS4wLyIKICAgIHhtbG5zOnBob3Rvc2hvcD0iaHR0cDovL25zLmFkb2JlLmNvbS9waG90b3Nob3AvMS4wLyIKICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIKICAgIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIgogICAgeG1sbnM6c3RFdnQ9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZUV2ZW50IyIKICAgdGlmZjpJbWFnZUxlbmd0aD0iMzIiCiAgIHRpZmY6SW1hZ2VXaWR0aD0iMzIiCiAgIHRpZmY6UmVzb2x1dGlvblVuaXQ9IjIiCiAgIHRpZmY6WFJlc29sdXRpb249IjEzOC4wIgogICB0aWZmOllSZXNvbHV0aW9uPSIxMzguMCIKICAgZXhpZjpQaXhlbFhEaW1lbnNpb249IjMyIgogICBleGlmOlBpeGVsWURpbWVuc2lvbj0iMzIiCiAgIGV4aWY6Q29sb3JTcGFjZT0iMSIKICAgcGhvdG9zaG9wOkNvbG9yTW9kZT0iMyIKICAgcGhvdG9zaG9wOklDQ1Byb2ZpbGU9InNSR0IgSUVDNjE5NjYtMi4xIgogICB4bXA6TW9kaWZ5RGF0ZT0iMjAyMi0w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
```
