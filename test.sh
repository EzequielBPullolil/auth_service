#!/bin/bash

# Definir las opciones y sus descripciones
OPTIONS=$(getopt -o l --long log -- "$@")
eval set -- "$OPTIONS"

# Inicializar la variable para el formato
format="pkgname"

# Leer las opciones
while true; do
    case "$1" in
    -l | --log)
        format="testname"
        shift
        ;;
    --)
        shift
        break
        ;;
    esac
done

# Ejecutar el comando gotestsum con las opciones
gotestsum --watch --format $format --hide-summary=output --rerun-fails
