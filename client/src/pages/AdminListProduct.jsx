import React from "react";

import Layout from "@/layouts/Default";

const AdminListProduct = () => {
	const title = "List Product";
	document.title = "WaysBeans | " + title;
	return (
		<Layout className='max-w-screen-lg mx-auto'>
			<>
				<section className='h-full mt-6 lg:mt-32 lg:px-6'>
					AdminListProduct
				</section>
			</>
		</Layout>
	);
};

export { AdminListProduct };
