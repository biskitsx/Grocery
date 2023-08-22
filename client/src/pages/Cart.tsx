import React, { useEffect } from 'react'
import Layout from '../components/Layout'
import axios from 'axios'
import CartProduct from '../components/CartProduct'
import { useDispatch, useSelector } from 'react-redux'
import { CartDTO, load } from '../store/cart'

function Cart() {
    const { cart } = useSelector((state) => state.cart) as CartDTO[]

    return (
        <Layout>
            <div className='flex flex-col gap-10'>
                {cart.map((product) =>
                    <CartProduct props={product} key={product.ID} />
                )}
            </div>
        </Layout>
    )
}

export default Cart