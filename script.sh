#!/bin/bash

backup_size=$(du -s "/home/a/Área de Trabalho/backup" | awk '{print $1}')

backup_size_gb=$(echo "scale=2; $backup_size / 1024^3" | bc)

memory_size=$(df -h / | awk '{print $4}')

echo "Espaço ocupado de Backup é: $backup_size_gb GB"
echo "Espaço de memória disponível na máquina é: $memory_size"
