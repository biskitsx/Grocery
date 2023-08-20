import './App.css'
import Layout from './components/Layout'
import { Route, Routes } from 'react-router-dom'
import All from './pages/All'

function App() {
  return (
    <div className='h-screen'>
        
        <Layout>
            <Routes>
                <Route path='/' element={<All/>}/>

            </Routes>
        </Layout>
        
    </div>
  )
}

export default App
