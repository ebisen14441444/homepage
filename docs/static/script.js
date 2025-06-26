const socket = new WebSocket("wss://homepage-w2x4.onrender.com/ws");
const tableBody = document.querySelector("#checkboxTable tbody");

socket.onmessage = (event) => {
  const data = JSON.parse(event.data);
  const cb = document.querySelector(`input[data-key="${data.key}"]`);
  if (cb) cb.checked = data.value;
};

for (let i = 1; i <= 9; i++) {
  const row = document.createElement("tr");
  const th = document.createElement("th");
  th.textContent = i;
  row.appendChild(th);

  ["pink", "green", "yellow", "blue"].forEach(color => {
    const td = document.createElement("td");
    const key = `${i}-${color}`;

// ラベル作成（判定範囲を広げる）
const label = document.createElement("label");
label.style.display = "inline-block";
label.style.padding = "10px";  
label.style.cursor = "pointer";

const cb = document.createElement("input");
cb.type = "checkbox";
cb.setAttribute("data-key", key);
cb.style.transform = "scale(1.0)";
cb.addEventListener("change", () => {
  socket.send(JSON.stringify({ key: key, value: cb.checked }));
});

label.appendChild(cb);
td.appendChild(label);

    row.appendChild(td);
  });

  tableBody.appendChild(row);
}

document.getElementById("resetBtn").addEventListener("click", () => {
  for (let i = 1; i <= 9; i++) {
    ["pink", "green", "yellow", "blue"].forEach(color => {
      const key = `${i}-${color}`;
      const cb = document.querySelector(`input[data-key="${key}"]`);
      if (cb) {
        cb.checked = false;
        socket.send(JSON.stringify({ key: key, value: false }));
      }
    });
  }
});

