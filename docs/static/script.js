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
    const cb = document.createElement("input");
    const key = `${i}-${color}`;
    cb.type = "checkbox";
    cb.setAttribute("data-key", key);
    cb.addEventListener("change", () => {
      socket.send(JSON.stringify({ key: key, value: cb.checked }));
    });
    td.appendChild(cb);
    row.appendChild(td);
  });

  tableBody.appendChild(row);
}
