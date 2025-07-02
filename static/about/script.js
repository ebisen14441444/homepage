const text=document.querySelector('#colorText');
const color = document.querySelector('#colorPicker');
const colorBg =()=>{
  text.textContent='カラーコード:${color.value}';
};

color.addEventListener('input',colorBg);

