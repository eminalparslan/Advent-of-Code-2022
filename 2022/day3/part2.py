totalPriority = 0

with open("input.txt") as f:
  lines = f.readlines()
  for i in range(0, len(lines), 3):
    s = (set(lines[i]) & set(lines[i+1]) & set(lines[i+2])) - {"\n"}
    c = ord(s.pop())
    if 65 <= c <= 90:
      totalPriority += c - 38
    elif 97 <= c <= 122:
      totalPriority += c - 96

print(f"{totalPriority=}")
