/* global document*/
// const CartClass = require('cart');
// class Codeshopping {
//     constructor(side, items) {
//       this.side = side;
//       this.items = items;
//     }
//     static writeItems() {
//       return new code(this.items);
//     }
//     get displayallCodeItem() {
//         return 1;
//     }
// }
const name = ['swedish', 'moscow', 'texas', 'picachu'];
const imgs = ['https://forum.golangbridge.org/uploads/default/original/2X/2/22304e0805368a2114fefbde7963bd750b97e950.png', 'http://photos3.meetupstatic.com/photos/event/a/0/2/highres_268082562.jpeg', 'http://photos3.meetupstatic.com/photos/event/9/7/d/6/600_434678870.jpeg', 'https://pbs.twimg.com/media/CnRZy5yXgAAgcD_.jpg'];
const imgs = ['swedish', 'moscow', 'texas', 'picachu'];

const price = [8, 8, 9, 10];
const sides = ['front', 'back'];


const createCode = (i) => {
  const o = { image : imgs[i], name : name[i], price :price[i], side :sides[0] };
  return o;
};

const codes = [];

for (let i = 0; i < imgs.length; i += 1) {
  codes.push(createCode(i));
}

console.log(codes);

// TODO: make a ajaxcall to the webapi to get the items

for (let i = 0; i < imgs.length; i += 1) {
  // TODO: make html of the items
  const tContainer = document.createElement('div');
  tContainer.className = 'col-sm-4 padding';
  tContainer.id = codes[i].name;
  document.getElementById('codes').appendChild(tContainer);

  const tImage = document.createElement('img');
  // tImage.className='img-responsive'
  tImage.src = codes[i].image;
  tImage.height = '300';
  tImage.style.borderRadius = '20px';
  tContainer.appendChild(tImage);

  const tName = document.createElement('h1');
  const nameTxt = document.createTextNode(codes[i].name);
  tName.appendChild(nameTxt);
  tContainer.appendChild(tName);

  const tCost = document.createElement('h4');
  const costTxt = document.createTextNode(codes[i].price);
  tCost.appendChild(costTxt);
  tContainer.appendChild(tCost);

  const tColor = document.createElement('h4');
  const colorTxt = document.createTextNode(codes[i].side);
  tColor.appendChild(colorTxt);
  tContainer.appendChild(tColor);

  const tButton = document.createElement('button');
  tButton.className = 'btn btn-primary btn-lg';
  tButton.innerHTML = 'Buy Now';
  tButton.style.color = 'blue';
  tButton.setAttribute('onclick', 'alert(\'hello\');');
  tContainer.appendChild(tButton);
}
