import React, { useState, useContext, useEffect } from "react";
import Header from "@/components/Header";
import { Navbar, Dropdown, Avatar, Button } from "flowbite-react";
import { Link, useNavigate } from "react-router-dom";

import { useQuery } from "react-query";
import { AppContext } from "@/context/AppContext";
import { API } from "@/lib/api";
import { setAuthToken } from "@/lib/api";

import Login from "@/components/Auth/Login";
import Register from "@/components/Auth/Register";

if (localStorage.token) {
	setAuthToken(localStorage.token);
}
export default function Layouts(props) {
	const [loginModal, setLoginModal] = useState(false);
	const [registerModal, setRegisterModal] = useState(false);
	const [state, dispatch] = useContext(AppContext);

	const navigate = useNavigate();

	const isLogout = () => {
		dispatch({
			type: "LOGOUT",
		});
		navigate("/");
		Toast.fire({
			icon: "success",
			title: "Logout Success, Bye👋",
		});
	};
	return (
		<>
			<div>
				<Header className={"fixed w-full shadow-lg z-30"}>
					<Navbar.Collapse className='p-1 !rounded-lg md:flex place-items-center'>
						{state.isLogin === true ? (
							<div className='flex md:flex-none justify-end pb-4 md:pb-0'>
								<Dropdown
									arrowIcon={false}
									inline={true}
									label={
										<Avatar
											alt='User settings'
											img={state.user.image}
											rounded={true}
										/>
									}
								>
									<Dropdown.Header>
										<span className='block text-sm font-bold'>
											{state.user.name}
										</span>
										<span className='block truncate text-sm font-medium'>
											{state.user.email}
										</span>
									</Dropdown.Header>

									{state.user.role === "admin" ? (
										<>
											<Link to={"/product_add"}>
												<Dropdown.Item className='hover:bg-coffee-100 hover:text-coffee-400 font-medium'>
													Add Product
												</Dropdown.Item>
											</Link>
											<Link to={"/product_list"}>
												<Dropdown.Item className='hover:bg-coffee-100 hover:text-coffee-400 font-medium'>
													List Product
												</Dropdown.Item>
											</Link>
										</>
									) : (
										<Link to={"/profile"}>
											<Dropdown.Item className='hover:bg-coffee-100 hover:text-coffee-400 font-medium'>
												Profile
											</Dropdown.Item>
										</Link>
									)}

									<Dropdown.Divider />
									<Dropdown.Item
										onClick={isLogout}
										className='hover:bg-rose-100 hover:text-rose-600 font-medium'
									>
										Logout
									</Dropdown.Item>
								</Dropdown>
							</div>
						) : (
							<div className='flex flex-col md:flex-row md:items-center gap-4 w-full'>
								<Button
									outline={true}
									color='light'
									onClick={() => setLoginModal(true)}
									className='hover:!bg-coffee-400 border-2 !border-coffee-400 text-coffee-400 hover:text-white px-6 py-0 font-semibold transition-all w-full md:w-auto'
								>
									Login
								</Button>
								<Button
									color='light'
									onClick={() => setRegisterModal(true)}
									className='!bg-coffee-400 hover:!bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 px-4 py-0 font-semibold transition-all w-full md:w-auto'
								>
									Register
								</Button>
							</div>
						)}
					</Navbar.Collapse>
				</Header>
				<div {...props}>{props.children}</div>
			</div>

			{/* Login */}
			<Login
				loginModal={loginModal}
				setLoginModal={setLoginModal}
				setRegisterModal={setRegisterModal}
			/>

			{/* Register */}
			<Register
				registerModal={registerModal}
				setRegisterModal={setRegisterModal}
				setLoginModal={setLoginModal}
			/>
		</>
	);
}
