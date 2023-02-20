import { useState, useContext } from "react";
import { Button, Modal, TextInput } from "flowbite-react";

import { useNavigate } from "react-router-dom";
import { useMutation } from "react-query";
import { AppContext } from "@/context/AppContext";
import { API } from "@/lib/api";
import Toast from "@/lib/alert";

const Login = ({ loginModal, setLoginModal, setRegisterModal }) => {
	const [state, dispatch] = useContext(AppContext);
	const [input, setInput] = useState({
		email: "",
		password: "",
	});
	let navigate = useNavigate();

	const handleInputChange = (e) => {
		setInput((prevState) => {
			return { ...prevState, [e.target.name]: e.target.value };
		});
	};

	const handleLogin = useMutation(async (e) => {
		e.preventDefault();
		try {
			// Configuration
			const config = {
				headers: {
					"Content-type": "application/json",
				},
			};
			// Data body
			const body = JSON.stringify(input);

			// Insert data for login process
			const response = await API.post("/login", body, config);

			// Checking process
			if (response.data != null) {
				// Send data to useContext
				dispatch({
					type: "SIGNIN",
					payload: response.data.data,
				});
				navigate("/");
				Toast.fire({
					icon: "success",
					title: "Successfully Sign In",
				});
			}
			e.target.reset();
			setLoginModal(false);
		} catch (err) {
			Toast.fire({
				icon: err.response.data.status,
				title: err.response.data.message,
			});
		}
	});

	return (
		<Modal
			size={"md"}
			show={loginModal}
			popup={true}
			dismissible={true}
			onClose={() => setLoginModal(false)}
		>
			<Modal.Header />
			<Modal.Body>
				<div className='space-y-6 px-2 pb-4'>
					<h3 className='text-3xl font-serif font-semibold text-black'>
						Login
					</h3>
					<form onSubmit={handleLogin.mutate} className='flex flex-col gap-4'>
						<div>
							<input
								type='email'
								name='email'
								placeholder='Email'
								onChange={handleInputChange}
								value={input.email}
								required
								className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
							/>
						</div>
						<div>
							<input
								type='password'
								name='password'
								placeholder='Password'
								onChange={handleInputChange}
								value={input.password}
								required
								className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
							/>
						</div>

						<Button
							type='submit'
							className='!bg-coffee-400 hover:!bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 font-semibold transition-all w-full'
						>
							Login
						</Button>
					</form>

					<div className='text-sm font-medium text-gray-500 dark:text-gray-300'>
						Don't have an account ? Klik{" "}
						<span
							onClick={() => {
								setLoginModal(false);
								setRegisterModal(true);
							}}
							className='text-black hover:underline hover:text-coffee-400 font-semibold cursor-pointer'
						>
							Here
						</span>
					</div>
				</div>
			</Modal.Body>
		</Modal>
	);
};
export default Login;
