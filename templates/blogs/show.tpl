{{ define "blogs/show.tpl" }}
    {{ template "layouts/header.tpl" .}}
        <h1 class="card-title">{{ .blog.Title }}</h1>
        <p class="card-text">{{ .blog.Body }}</p>
    {{ template "layouts/footer.tpl" .}}
{{ end }}