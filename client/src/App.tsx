import './App.css';
import Layout from './components/Layout';
import { Navigate, Route, Routes } from 'react-router-dom';
import All from './pages/All';
import Login from './pages/Login';
import { useDispatch, useSelector} from 'react-redux';
import { login } from './store/user';
import { useEffect } from 'react';
import Cart from './pages/Cart';

function App() {
  const { user } = useSelector((state) => state.user);
  const dispatch = useDispatch();

  useEffect(() => {
    const userInfo = localStorage.getItem('user');
    const objUser = JSON.parse(userInfo);
    if (objUser) {
      dispatch(login(objUser));
    }
  }, []);

  return (
    <div className='h-screen'>
      <Routes>
        <Route path='/' name='home' element={user ? <All /> : <Navigate to="/login"/>} />
        <Route path='/login' name='login' element={<Login />} />
        <Route path='/cart' name='login' element={<Cart />} />
      </Routes>
    </div>
  );
}

export default App;
