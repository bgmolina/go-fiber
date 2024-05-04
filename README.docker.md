<a href="README.md">
  <img
    align="right"
    src="https://img.shields.io/badge/Inicio-161b22?style=for-the-badge&logoColor=white&logo=github"
    alt="Inicio"
  />
</a>

# Go fiber API 🔄

<div>
  <a href="https://go.dev/" target="_blank">
    <img
      src="https://img.shields.io/badge/v1.22-gray?style=flat&logo=go&logoColor=white&label=Golang&labelColor=79D4FD"
      alt="Golang"
    />
  </a>
  <a href="https://gofiber.io/" target="_blank">
    <img
      src="https://img.shields.io/badge/v2.52.4-gray?style=flat&logo=fiber&label=Fiber&labelColor=00acd7"
      alt="Fiber"
    />
  </a>
</div>

### Requisitos previos 📝

- Docker Compose **versión 2.x**
- Ejecutar bash desde la carpeta **docker**

#### Puertos habilitados

- [**3000**](http://localhost:3000/docs) para desarrollo
- [**4000**](http://localhost:4000/docs) para producción

#### Menu de opciones 📋

```bash
$ bash deployment.sh
```

```bash
====================
 Go Fiber deployment
====================
1) Deploy on development mode 🛠
2) Delete on development mode 🗑️
3) Deploy on production mode 🚀
4) Delete on production mode 🗑️
5) Quit 👋
Select an option and press Enter 👆:
```

### Demo 🎬

<img width="500" src="./demo/docker.gif"/>

## Información relevante 📑

### Imagen

```bash
$ docker images
```

| REPOSITORY | TAG             | SIZE   |
| ---------- | --------------- | ------ |
| golang     | 1.22-alpine3.19 | 230MB  |
| go-fiber   | 1.0.0           | ~255MB |

### Contenedor

```bash
$ docker ps
```

| NAME          | PORTS                  |
| ------------- | ---------------------- |
| go-fiber      | 0.0.0.0:3000->3000/tcp |
| go-fiber-prod | 0.0.0.0:4000->4000/tcp |
