import React from "react";
import Layout from "@/layouts/Default";
import moment from "moment";
import { useState } from "react";

const Profile = () => {
	const [showMore, setShowMore] = useState(null);
	console.log(showMore);
	const showMoreClick = (n) => {
		if (showMore === n) {
			setShowMore(null);
		} else {
			setShowMore(n);
		}
	};
	const coffeeData = [
		{
			img: "/uploads/product1.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			qty: "2",
		},
		{
			img: "/uploads/product2.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			qty: "2",
		},
		{
			img: "/uploads/product3.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			qty: "2",
		},
		{
			img: "/uploads/product4.png",
			name: "RWANDA Beans",
			price: "Rp. 299.000",
			qty: "2",
		},
	];
	const title = "Profile";
	document.title = "WaysBeans | " + title;

	return (
		<Layout className='max-w-screen-xl mx-auto'>
			<section className='mt-20 mx-16'>
				<div className='flex gap-8 justify-between'>
					{/* My Profile */}
					<div className=''>
						<h3 className='text-2xl font-semibold'> My Profile</h3>
						<div className='flex mt-7'>
							<div className='h-56 w-44 rounded-lg overflow-hidden'>
								<img
									src='https://api.dicebear.com/5.x/shapes/svg?seed=hyujisf'
									className='object-fit object-cover w-full h-full'
									alt=''
								/>
							</div>
							<div className='mx-4 flex flex-col gap-5'>
								<div className=''>
									<h4 className='font-medium text-xl'>Fullname</h4>
									<p>Wahyudi Chridianto</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Email</h4>
									<p>hyujisf@mail.com</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Phone</h4>
									<p>08123456789</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Post Code</h4>
									<p>112</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Address</h4>
									<p>Disana</p>
								</div>
							</div>
						</div>
					</div>
					<div className=' w-[38rem]'>
						<h4 className='text-2xl font-semibold'>My Transaction</h4>

						<div className='flex flex-col gap-4 mt-7'>
							{coffeeData.map((data, k) => (
								<div key={k} className='bg-coffee-100 py-3 px-6'>
									<h4 className='font-semibold'>{coffeeData[0].name}</h4>{" "}
									<span
										onClick={() => showMoreClick(k)}
										className='cursor-pointer'
									>
										show more
									</span>
									<br />
									<span>
										<b>{moment(Date.now()).format("dddd")}</b>,
										{moment(Date.now()).format("DD MMMM yyyy")}
									</span>
									<div className={showMore == k ? "flex" : "hidden"}>
										{coffeeData.slice(1).map((data, k) => (
											<div key={k} className='flex gap-4'>
												<div>
													<img src={data.img} width={100} alt='' />
												</div>
												<div className=''>
													<h4 className='font-semibold text-coffee-400'>
														{data.name}
													</h4>
												</div>
											</div>
										))}
									</div>
								</div>
							))}
						</div>
					</div>
				</div>
			</section>
		</Layout>
	);
};

export { Profile };
