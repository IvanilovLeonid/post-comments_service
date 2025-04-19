#!/bin/bash


# Убедитесь, что выполняете скрипт из корневой директории проекта
PROJECT_ROOT=$(pwd)
MOCKS_DIR="$PROJECT_ROOT/test/mocks"

# Генерация моков для репозиториев
mockgen -source="$PROJECT_ROOT/internal/core/repository/repository.go" \
        -destination="$MOCKS_DIR/post_repository.go" \
        -package=mocks

# Генерация моков для сервисов
mockgen -source="$PROJECT_ROOT/internal/core/ports/service.go" \
        -destination="$MOCKS_DIR/post_service.go" \
        -package=mocks