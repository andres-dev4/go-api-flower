# go-api-flower
api simple in go to lab with gitops

# 🧪 Laboratorio DevOps con Go, Docker, Swagger y GitOps

Este proyecto es una API en Go que devuelve datos de flores y saluda personas. Está pensado para adquirir experiencia práctica en:

- Desarrollo backend en Go
- Contenerización con Docker
- Documentación Swagger/OpenAPI
- CI/CD con GitHub Actions
- GitOps con Argo CD o Flux CD
- Kubernetes y observabilidad

---

## ✅ Fase 1: API sólida y documentada

- [x] Crear API en Go (GET /flowers y POST /greet)
- [x] Documentar con Swagger (`swaggo/swag`)
- [x] Crear `Dockerfile` y `docker-compose.yml`
- [ ] Agregar logs estructurados (`logrus` o `zap`)
- [ ] Agregar validaciones de input (`binding`)
- [ ] Agregar health checks (`/healthz`)
- [ ] Agregar pruebas unitarias e integración

---

## 🚀 Fase 2: GitOps y CI/CD

- [ ] Subir proyecto a GitHub
- [ ] Crear acción de CI (lint, test, build)
- [ ] Crear acción de CD (build + push a DockerHub o GHCR)
- [ ] Crear manifiestos Kubernetes:
  - `deployment.yaml`
  - `service.yaml`
- [ ] Crear estructura GitOps con Kustomize (`dev`, `prod`)
- [ ] Instalar Argo CD en Minikube
- [ ] Hacer primer despliegue GitOps

---

## 📈 Fase 3: Observabilidad

- [ ] Agregar Prometheus + Grafana en Minikube
- [ ] Exponer métricas en `/metrics` (`promhttp`)
- [ ] Agregar tracing con OpenTelemetry
- [ ] Centralizar logs (Grafana Loki)

---

## 🔐 Fase 4: Casos reales

- [ ] Agregar autenticación con JWT
- [ ] Agregar middlewares de autorización y logging
- [ ] Agregar SPA en React que consuma la API
- [ ] Simular errores y caídas (resiliencia)
- [ ] Agregar pruebas de carga (K6)

---

## 🧠 Aprendizajes esperados

- Buenas prácticas en Go para backend
- Dockerfile y Compose para desarrollo
- CI/CD con GitHub Actions
- GitOps puro con Kubernetes
- Observabilidad y resiliencia
