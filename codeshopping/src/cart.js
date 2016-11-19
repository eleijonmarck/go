export class Cart {

  constructor(cartitems) {
    if ( cartitems.length === 0) {
        this.cartitems = [];
    }
    else {
        this.cartitems = cartitems;
    }
  }

  get getAllItems() {
    return this.cartitems;
  }

  set addItemToCart(cartitem) {
    this.cartitems.push(cartitem);
  }
}

