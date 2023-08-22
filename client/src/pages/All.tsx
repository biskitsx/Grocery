
import Product from '../components/Product'
import Layout from '../components/Layout'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { useDispatch } from 'react-redux'
import { load } from '../store/cart'

interface Product {
    name: string
    price: number
    picture: string
    ID: number
}
function All() {
    const [product, setProduct] = useState<Product[]>([])
    const dispatch = useDispatch()
    useEffect(() => {
        const fetchData = async () => {
            const res = await axios.get("http://localhost:3001/api/product")
            const data: Product[] = res.data as Product[]
            setProduct(data)
        }
        fetchData()
    }, [])
    return (
        <Layout>
            <div className='grid grid-cols-3 gap-10'>
                {product.map((item) =>
                    <Product props={item} key={item.ID} />
                )}
            </div>
        </Layout>
    )
}

export default All