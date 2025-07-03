const btn = document.querySelector('#btn');

btn.addEventListener('click',() =>{
  document.body.classList.toggle('dark-theme');

  if (btn.textContent==='ダークモード'){
    btn.textContent = "ライトモード";
  } else {
    btn.textContent = "ダークモード";
  }
});