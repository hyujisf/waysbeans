import React from "react";
import Layout from "@/layouts/Default";
import moment from "moment";
import { useState } from "react";
import { useQuery } from "react-query";
import { API } from "@/lib/api";
import logo from "/img/Waysbeans.svg";
import qrCode from "/img/qrcode.svg";
import { toCurrency } from "@/lib/currency";

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
							{transactions?.length != 0 ? (
								<>
									{transactions?.map((items, k) => (
										<div key={k} className='bg-coffee-100 py-3 px-6'>
											{/* <h4 className='font-semibold'>{data.product[0].name}</h4>{" "} */}
											<span
												onClick={() => showMoreClick(k)}
												className='cursor-pointer'
											>
												show more
											</span>
											<br />
											{items?.product?.map((data, idx) => (
												<div className='' key={idx}>
													<span>
														{/* <b>{moment(Date.now()).format("dddd")}</b>,
														{moment(Date.now()).format("DD MMMM yyyy")} */}
													</span>

													<div>
														<div className='mb-3'>
															<div>
																<img
																	src={data?.product?.image}
																	alt='fotokopi'
																	style={{ width: 100, borderRadius: 5 }}
																/>
															</div>
															<div>
																<div>
																	<h5>{data?.product?.name}</h5>
																	<p>
																		<b>
																			{moment(items?.created_at).format("dddd")}
																		</b>
																		<span>
																			{moment(items.created_at).format(
																				"DD MMMM yyyy"
																			)}
																		</span>
																	</p>
																	<p>Qty: {data?.qty}</p>
																</div>
																<div className='mt-1' style={{ fontSize: 15 }}>
																	<p className='my-1'>
																		Price : {toCurrency(data?.subtotal)}
																	</p>
																</div>
															</div>
														</div>
													</div>
												</div>
											))}
											<div key={k} className='flex gap-4'>
												<img className='w-50' src={logo} alt='' />
												<br />
												<br />
												<img src={qrCode} alt='' />
												<div
													className='text-center w-75 m-auto my-3 fw-semibold'
													style={{
														backgroundColor: "rgba(0, 209, 255, .3)",
														color: "#34a8eb",
													}}
												>
													{items?.status}
												</div>
												<div className='text-center w-75 m-auto my-3 fw-normal'>
													Subtotal:{toCurrency(items?.total)}
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
