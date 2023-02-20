import { useState, useContext } from "react";
import { Button, Modal, TextInput } from "flowbite-react";

import { useMutation } from "react-query";
import { AppContext } from "@/context/AppContext";
import { API } from "@/lib/api";
import Toast from "@/lib/alert";

const Register = ({ registerModal, setRegisterModal, setLoginModal }) => {
	const [input, setInput] = useState({
		name: "",
		email: "",
		password: "",
	});

	const handleInputChange = (e) => {
		setInput((prevState) => {
			return { ...prevState, [e.target.name]: e.target.value };
		});
	};

	const handleRegister = useMutation(async (e) => {
		try {
			e.preventDefault();

			// Configuration Content-type
			const config = {
				headers: {
					"Content-type": "application/json",
				},
			};

			// Data body
			const body = JSON.stringify(input);

			// Insert data user to database
			const response = await API.post("/register", body, config);

			// Notification
			if (response.data != null) {
				setInput({
					name: "",
					email: "",
					password: "",
				});

				Toast.fire({
					icon: "success",
					title: "Successfully Sign Up",
				});
			} else {
				Toast.fire({
					icon: response.data.status,
					title: response.data.message,
				});
			}
			e.target.reset();
			setRegisterModal(false);
			setLoginModal(true);
		} catch (err) {
			Toast.fire({
				icon: err.response.data.status,
				title: err.response.data.message,
			});
		}
	});

	return (
		<Modal
			className='bottom-0 '
			size={"md"}
			show={registerModal}
			popup={true}
			dismissible={true}
			onClose={() => setRegisterModal(false)}
		>
			<Modal.Header />
			<Modal.Body>
				<div className='space-y-6 px-2 pb-4'>
					<h3 className='text-3xl font-serif font-semibold text-black'>
						Register
					</h3>
					<form
						onSubmit={handleRegister.mutate}
						className='flex flex-col gap-4'
					>
						<div>
							<TextInput
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
							<TextInput
								type='password'
								name='password'
								placeholder='Password'
								onChange={handleInputChange}
								value={input.password}
								required
								className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
							/>
						</div>
						<div>
							<TextInput
								type='text'
								name='name'
								placeholder='Fullname'
								onChange={handleInputChange}
								value={input.name}
								required
								className='w-full rounded-lg bg-coffee-400/25 placeholder:text-coffee-300 text-coffee-400 border-2 border-coffee-400'
							/>
						</div>
						<Button
							type='submit'
							className='!bg-coffee-400 hover:!bg-transparent border-2 !border-coffee-400 text-white hover:text-coffee-400 font-semibold transition-all w-full'
						>
							Register
						</Button>
					</form>
					<div className='text-sm font-medium text-gray-500 dark:text-gray-300'>
						Already have an account? klik{" "}
						<span
							onClick={() => {
								setRegisterModal(false);
								setLoginModal(true);
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
export default Register;
