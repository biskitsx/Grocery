import React, { useEffect } from 'react'
import Layout from '../components/Layout'
import axios from 'axios'

function Cart() {
    const getCart = async () => {
        const res = await axios.get("http://localhost:3001/api/user/cart")
        console.log(res.data)
    }
    useEffect(() => {
        getCart()
    }, [])
    return (
        <Layout>
            <div className='flex flex-col'>
                <div className="alert bg-base-100 h-32 flex shadow-md">
                    <div className='h-full shadow-md'>
                        <img src="/hbg.webp" alt="hbg" className='h-full' />
                    </div>
                    <span>Hamburger</span>
                    <div className=''>
                        <button className="btn btn-sm btn-primary">Accept</button>
                    </div>
                </div>
            </div>

        </Layout>
    )
}

export default Cart