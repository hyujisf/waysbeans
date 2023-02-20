import { useEffect, useContext } from "react";
import { Routes, Route } from "react-router-dom";

import { API, setAuthToken } from "@/lib/api";
import { AppContext } from "@/context/AppContext";

// Pages
import * as Pages from "@/pages/index";

if (localStorage.token) {
	setAuthToken(localStorage.token);
}
function App() {
	const [state, dispatch] = useContext(AppContext);

	useEffect(() => {
		if (localStorage.token) {
			setAuthToken(localStorage.token);
		}
	}, [state]);
	// let navigate = useNavigate();
	// console.clear();
	console.log(state);

	const checkAuth = async () => {
		try {
			const response = await API.get("/check-auth");

			// If the token incorrect
			if (response.data.code === 404) {
				return dispatch({
					type: "AUTH_ERROR",
				});
			}

			// Get user data
			let payload = response.data.data;
			// Get token from local storage
			payload.token = localStorage.token;

			// Send data to useContext
			dispatch({
				type: "SIGNIN",
				payload,
			});
		} catch (error) {
			console.log(error);
		}
	};

	useEffect(() => {
		if (localStorage.token) {
			checkAuth();
		}
	}, []);
	return (
		<Routes>
			<Route path='/' element={<Pages.Home />} />
			<Route path='/detail/:id' element={<Pages.Detail />} />
			<Route path='/profile' element={<Pages.Profile />} />
			<Route path='/cart' element={<Pages.CartProduct />} />

			<Route path='/' element={<Pages.AdminDashboard />} />
			<Route path='/product_add' element={<Pages.AdminAddProduct />} />
			<Route path='/product_list' element={<Pages.AdminListProduct />} />
		</Routes>
	);
}

export default App;
