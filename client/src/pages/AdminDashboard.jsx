import React from "react";

import Layout from "@/layouts/Default";

const AdminDashboard = () => {
	const title = "Dashboard";
	document.title = "WaysBeans | " + title;
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<>
				<section className='h-full mt-6 lg:mt-32 lg:px-6'>Admin Page</section>
			</>
		</Layout>
	);
};

export { AdminDashboard };
