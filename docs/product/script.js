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
    cb.addEventListener("change", () => {
    updateCheck(i, color, cb.checked);
  });
    td.appendChild(cb);
    tr.appendChild(td);
  }

  tbody.appendChild(tr);
}

// POST（変更時に呼ぶ）
function updateCheck(row, color, value) {
  fetch("http://localhost:8080/api/checks", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ row, color, value })
  });
}

// GET（初回読み込み時に呼ぶ）
async function loadChecks() {
  const res = await fetch("http://localhost:8080/api/checks");
  const checks = await res.json();
  for (const [key, checked] of Object.entries(checks)) {
    const cb = document.querySelector(`input[data-key="${key}"]`);
    if (cb) cb.checked = checked;
  }
}


