A mini RSVP web app built as a learning project while working through Pro Go: The Complete Guide to Programming Reliable and Efficient Software Using Golang (2022) by Adam Freeman.

## Features / What I Learned

- **Go Templates (`html/template`)**  
  Dynamically render HTML pages with data from Go structs. Explored `{{ . }}` and named blocks like `{{ block "body" . }}` and `{{ block "footer" . }}` for modular layouts.

- **HTTP Routing and Handlers**  
  Built simple routes (`/`, `/form`, `/list`) using `http.HandleFunc` to navigate pages and handle GET/POST requests.

- **Structs and Pointers**  
  Used Go structs (`Rsvp`) with pointers to efficiently pass and update data across templates and handlers.

- **Form Validation and Error Handling**  
  - Required fields (`Name`, `Email`, `Phone`)  
  - Valid email format (`strings.Contains`)  
  - Phone number must be numeric and minimum 11 digits (`unicode.IsDigit`, `len()`)  
  Displayed errors back in the form using a `formData` struct.

- **Template Blocks for Modular Layout**  
  Used `{{ define "body" }}` and `{{ define "footer" }}` to separate content and make templates reusable.
