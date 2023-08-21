import React, { useState } from 'react'
import Nav from '../components/Nav'
import axios from 'axios'
import { useDispatch } from 'react-redux'
import { login } from '../store/user'
import { useNavigate } from 'react-router-dom'

function Login() {
    const [username,setUsername] = useState("")
    const [password,setPassword] = useState("")
    const dispatch = useDispatch()
    const navigate = useNavigate()
    const handleButton = async (e: any) => {
        try {
            const response = await axios.post("http://localhost:3001/api/auth/login", {username, password})
            const data = response.data
            dispatch(login(data))
            localStorage.setItem("user", JSON.stringify(data))
            navigate('/')
        } catch (error: any) {
            console.log(error.response.data)
        }
    }
    return (
        <div className="hero min-h-screen bg-base-200">
            <div className="hero-content flex-col lg:flex-row-reverse">
                <div className="text-center lg:text-left">
                    <h1 className="text-5xl font-bold">Login now!</h1>
                    {/* <p className="py-6"></p> */}
                </div>
                <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
                    <div className="card-body">
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Email</span>
                            </label>
                            <input type="text" placeholder="email" className="input input-bordered" value={username} onChange={(e)=>setUsername(e.target.value)}/>
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Password</span>
                            </label>
                            <input type="password" placeholder="password" className="input input-bordered" value={password} onChange={(e)=>setPassword(e.target.value)}/>
                            <label className="label">
                                <a href="#" className="label-text-alt link link-hover">Forgot password?</a>
                            </label>
                        </div>
                        <div className="form-control mt-6">
                            <button className="btn btn-primary" onClick={handleButton}>Login</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    )
}

export default Login
