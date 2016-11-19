'use strict';

describe('cart', () => {
    const addtocart;

    beforeEach(function(_addtocart_) {
        addtocart = _addtocart_;
    }));
    describe("add one item to cart", () => {
        it('should add the item to the cart', addtocart(item) => {
            expect(cart.getallitems,item);
        });
