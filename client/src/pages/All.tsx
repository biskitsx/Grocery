
import Product from '../components/Product'
import Layout from '../components/Layout'



function All() {
    return (
        <Layout>
            <div className='grid grid-cols-3 gap-10'>
                <Product />
                <Product />
                <Product />
            </div>
        </Layout>
    )
}

export default All