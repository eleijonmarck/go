export class cartitem {
    
    constructor(image, name, price, side, quantity) {
      this.image = image;
      this.name = name;
      this.price = price;
      this.side = side;
      this.quantity = quantity;
    }

    get getItem() {
      return { image : this.image; name : this.name; price : this.price; side : this.side }
    }
}

const items = [{
  "id": "28dca95f-c730-442f-996b-c843d049cf2e",
  "name": "aenean lectus pellentesque eget nunc donec quis orci eget orci vehicula condimentum curabitur in libero ut massa volutpat",
  "image": "http://dummyimage.com/242x121.png/dddddd/000000",
  "price": 35,
  "quantity": 1
}, {
  "id": "96368dd7-a5c6-4ce0-ae5e-5df35da9d90c",
  "name": "dolor sit amet consectetuer adipiscing elit proin risus praesent lectus",
  "image": "http://dummyimage.com/156x244.bmp/5fa2dd/ffffff",
  "price": 75,
  "quantity": 2
}, {
  "id": "4498d245-7c3a-4287-b5fc-293fbde71bfe",
  "name": "eget eros elementum pellentesque quisque porta volutpat erat quisque erat eros viverra eget",
  "image": "http://dummyimage.com/101x118.png/5fa2dd/ffffff",
  "price": 96,
  "quantity": 3
}, {
  "id": "52cc971c-300c-435e-a094-d0bbc49d14e8",
  "name": "diam neque vestibulum eget vulputate ut ultrices vel augue vestibulum ante ipsum primis in faucibus orci luctus et ultrices",
  "image": "http://dummyimage.com/107x240.jpg/ff4444/ffffff",
  "price": 6,
  "quantity": 4
}, {
  "id": "e882de2f-c32d-4eb2-beb7-678edfd40cf0",
  "name": "ut suscipit a feugiat et eros vestibulum ac est lacinia nisi venenatis tristique fusce congue diam id",
  "image": "http://dummyimage.com/215x209.bmp/dddddd/000000",
  "price": 17,
  "quantity": 5
}, {
  "id": "95e49873-ccaa-444e-b1b2-16b5bd85ccdb",
  "name": "sed tristique in tempus sit amet sem fusce consequat nulla nisl nunc nisl duis bibendum felis sed",
  "image": "http://dummyimage.com/134x139.png/cc0000/ffffff",
  "price": 58,
  "quantity": 6
}, {
  "id": "1e608045-0c22-4a0a-b5e7-000c91940579",
  "name": "at feugiat non pretium quis lectus suspendisse potenti in eleifend quam a odio in hac",
  "image": "http://dummyimage.com/151x242.jpg/dddddd/000000",
  "price": 49,
  "quantity": 7
}, {
  "id": "993de736-24f8-40ce-a605-683748cc5574",
  "name": "posuere cubilia curae duis faucibus accumsan odio curabitur convallis duis consequat dui nec nisi volutpat eleifend donec ut dolor morbi",
  "image": "http://dummyimage.com/164x189.bmp/5fa2dd/ffffff",
  "price": 73,
  "quantity": 8
}, {
  "id": "df31e690-d149-4882-b40b-477a38f4d3ff",
  "name": "ultrices posuere cubilia curae mauris viverra diam vitae quam suspendisse",
  "image": "http://dummyimage.com/216x106.jpg/5fa2dd/ffffff",
  "price": 97,
  "quantity": 9
}, {
  "id": "7900c96d-a5dc-4640-98d8-3e8b75f0b2af",
  "name": "feugiat non pretium quis lectus suspendisse potenti in eleifend quam",
  "image": "http://dummyimage.com/246x206.bmp/ff4444/ffffff",
  "price": 4,
  "quantity": 10
}];
