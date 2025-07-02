const text=document.querySelector('#colorText');
const color = document.querySelector('#colorPicker');
const colorBg =()=>{
  document.body.style.backgroundColor = color.value;
  text.textContent=`カラーコード:${color.value}`;
};

color.addEventListener('input',colorBg);

