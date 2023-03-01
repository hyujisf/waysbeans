import React, { useState, useContext } from "react";
import Layout from "@/layouts/Default";
import moment from "moment";
import { useQuery } from "react-query";
import { API } from "@/lib/api";
import logo from "/img/Waysbeans.svg";
import qrCode from "/img/qrcode.svg";
import { toCurrency } from "@/lib/currency";
import { AppContext } from "@/context/AppContext";

const Profile = () => {
	const [showMore, setShowMore] = useState(null);
	const [state, dispatch] = useContext(AppContext);
	// console.log(showMore);
	const showMoreClick = (n) => {
		if (showMore === n) {
			setShowMore(null);
		} else {
			setShowMore(n);
		}
	};
	let { data: transactions } = useQuery("transactionsCache", async () => {
		const response = await API.get("/user-transaction");
		return response.data.data;
	});
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
									src={state.user.image}
									className='object-fit object-cover w-full h-full'
									alt=''
								/>
							</div>
							<div className='mx-4 flex flex-col gap-5'>
								<div className=''>
									<h4 className='font-medium text-xl'>Fullname</h4>
									<p>{state.user.name}</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Email</h4>
									<p>{state.user.email}</p>
								</div>
								{/* <div className=''>
									<h4 className='font-medium text-xl'>Phone</h4>
									<p>{state.user.phone || "-"}</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Post Code</h4>
									<p>{state.user.postcode || "-"}</p>
								</div>
								<div className=''>
									<h4 className='font-medium text-xl'>Address</h4>
									<p>{state.user.address || "-"}</p>
								</div> */}
							</div>
						</div>
					</div>
					<div className=' w-[38rem]'>
						<h4 className='text-2xl font-semibold'>My Transaction</h4>

						<div className='flex flex-col gap-4 mt-7'>
							{transactions?.length != 0 ? (
								<>
									{transactions?.map((items, k) => (
										<div key={k} className='bg-coffee-100 p-4'>
											<div className='flex w-full gap-4'>
												{items?.product?.map((data, idx) => (
													<div className='flex w-full' key={idx}>
														<div>
															<img
																src={data?.product?.image}
																alt={data?.product?.name}
																className='h-[10rem]'
															/>
														</div>
														<div>
															<div>
																<h5 className='font-bold text-xl'>
																	{data?.product?.name}
																</h5>
																<p>
																	<b>
																		{moment(items?.created_at).format("dddd")}
																	</b>
																	,{" "}
																	<span>
																		{moment(items.created_at).format(
																			"DD MMMM yyyy"
																		)}
																	</span>
																</p>
															</div>
															<div>
																<p className='my-1'>
																	Price : {toCurrency(data?.subtotal)}
																</p>
																<p>Qty: {data?.qty}</p>

																<p>Subtotal:{toCurrency(items?.total)}</p>
															</div>
														</div>
													</div>
												))}
												<div className='flex flex-col gap-4 justify-items-center w-72'>
													<img className='w-32 mx-auto' src={logo} alt='' />
													<img src={qrCode} alt='' className='w-20 mx-auto' />
													<div
														className={`font-medium capitalize text-center py-2 w-full rounded-lg ${
															items?.status === "success"
																? "bg-lime-200 text-lime-600"
																: items.status === "failed"
																? "bg-red-300 text-red-600"
																: items.status === "pending"
																? "bg-amber-300 text-orange-400"
																: ""
														}`}
													>
														{items?.status}
													</div>
												</div>
											</div>
										</div>
									))}
								</>
							) : (
								<div className='text-center '>
									<div className=''>No data Transactions, Let's Shopping</div>
								</div>
							)}
						</div>
					</div>
				</div>
			</section>
		</Layout>
	);
};

export { Profile };
