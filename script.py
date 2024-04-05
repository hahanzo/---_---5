import json

# Зчитування даних з файлу
with open('data.txt', 'r') as file:
    data = file.readlines()

# Розділення рядків на список чисел, ігноруючи недійсні значення
distances = []
for line in data:
    values = line.strip().split()
    numeric_values = []
    for value in values:
        if value.isdigit():  # Перевірка, чи є значення числом
            numeric_values.append(int(value))
        else:
            numeric_values.append(0)
    distances.append(numeric_values)

cities = ["Вінниця", "Дніпро", "Донецьк", "Житомир", "Запоріжжя", "Ів-Франківськ", "Київ", "Кропивницький", "Луганськ", "Луцьк", "Львів", "Миколаїв", "Одеса", "Полтава", "Рівне", "Сімферополь", "Суми", "Тернопіль", "Ужгород", "Харків", "Херсон", "Хмельницький", "Черкаси", "Чернівці", "Чернігів"]

# Створення порожнього словника для зберігання даних
city_data = []

# Заповнення даними
for i in range(len(cities)):
    city_distances = []
    for j in range(len(cities)):
        if i != j:  # Перевірка на наявність даних та ігнорування нульових значень
            city_distances.append({"city": cities[j], "distance": distances[i][j]})
    city_data.append({"name": cities[i], "connections": city_distances})

# Запис даних у JSON файл
with open('distances.json', 'w' ,encoding='utf-8') as outfile:
    json.dump(city_data, outfile, indent=4, ensure_ascii=False)
