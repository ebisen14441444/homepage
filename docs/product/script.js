// 4×9のチェック表を作る例
const tbody = document.querySelector("#checkboxTable tbody");

for (let i = 1; i <= 9; i++) {
  const tr = document.createElement("tr");

  const th = document.createElement("th");
  th.textContent = i;
  tr.appendChild(th);

  for (const color of ["pink", "green", "yellow", "blue"]) {
    const td = document.createElement("td");
    const cb = document.createElement("input");
    cb.type = "checkbox";
    cb.name = `row${i}-${color}`;
    td.appendChild(cb);
    tr.appendChild(td);
  }

  tbody.appendChild(tr);
}
