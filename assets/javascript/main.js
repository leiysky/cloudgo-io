window.onload = function() {
  const addButton = document.getElementById('addButton');
  addButton.onclick = () => {
    const li = document.createElement('li');
    li.setAttribute('class', 'item');
    li.textContent = 'hello world';
    document.getElementsByClassName('list')[0].appendChild(li);
  }

  const removeButton = document.getElementById('removeButton');
  removeButton.onclick = () => {
    const li = document.getElementsByClassName('list')[0].firstChild;
    document.getElementsByClassName('list')[0].removeChild(li);
  }
}