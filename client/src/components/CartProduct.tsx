import React from 'react'

function CartProduct({ props }) {
    console.log(props)
    return (
        <div className="alert bg-base-100 h-32 flex shadow-md justify-between">
            <div className='h-28 shadow-md w-28'>
                <img src={props.picture} alt="hbg" className='object-cover w-full h-full' />
            </div>
            <span>{props.name}</span>
            <div className=''>
                <button className="btn btn-sm btn-primary">Delete</button>
            </div>
        </div>
    )
}

export default CartProduct