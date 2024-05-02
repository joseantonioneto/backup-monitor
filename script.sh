#!/bin/bash

backup_size=$(du -s "/home/a/Área de Trabalho/backup" | awk '{print $1}')

backup_size_gb=$(echo "scale=2; $backup_size / 1024^3" | bc)

memory_size=$(df -h / | awk '/\// {print $4}')

percent_use=$(df -h / | awk '/\// {print $5}')

present_date=$(date +"%Y-%m-%d")

echo "Espaço ocupado de Backup: $backup_size_gb GB"
echo "Espaço disponível: $memory_size"
echo "Uso de espaço de disco: $percent_use"
echo "Data atual: $present_date"
