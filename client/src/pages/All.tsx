import React from 'react'
import Product from '../components/Product'



function All() {
  return (
        <div className='grid grid-cols-3 gap-10 px-10'>
            <Product/>
            <Product/>
            <Product/>
        </div>
  )
}

export default All