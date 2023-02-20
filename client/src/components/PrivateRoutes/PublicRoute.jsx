import { useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { MyContext } from "../../store/Store";

const PublicRoute = () => {
	const { loginState } = useContext(MyContext);

	return (
		<>
			{loginState?.isLogin && loginState.userInfo?.role === "admin" ? (
				<Navigate to='/transactions' />
			) : (
				<Outlet />
			)}
		</>
	);
};
export default PublicRoute;
