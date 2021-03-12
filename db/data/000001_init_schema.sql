

--
-- Data for Name: CTHEA_ENDPOINT; Type: TABLE DATA; Schema: manager; Owner: postgres
--

COPY manager."CTHEA_ENDPOINT" (id_endpoint, name, description, creation_date, edit_date, status) FROM stdin;
1	alfa-blephaclean	\N	2021-03-10 04:25:34.740338	\N	t
2	alfabeta-blephacelan	\N	2021-03-10 04:25:34.740338	\N	t
3	beta-blephaclean	\N	2021-03-10 04:25:34.740338	\N	t
4	blephaclean	\N	2021-03-10 04:25:34.740338	\N	t
5	gamma-blephaclean	\N	2021-03-10 04:25:34.740338	\N	t
6	hyabak-alfa	\N	2021-03-10 04:25:34.740338	\N	t
7	hyabak-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
8	hyabak-beta	\N	2021-03-10 04:25:34.740338	\N	t
9	hyabak	\N	2021-03-10 04:25:34.740338	\N	t
10	index	\N	2021-03-10 04:25:34.740338	\N	t
11	inicio	\N	2021-03-10 04:25:34.740338	\N	t
12	monolatan-alfa	\N	2021-03-10 04:25:34.740338	\N	t
13	monolatan-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
14	monolatan-beta	\N	2021-03-10 04:25:34.740338	\N	t
15	monolatan	\N	2021-03-10 04:25:34.740338	\N	t
16	monolatangamma	\N	2021-03-10 04:25:34.740338	\N	t
17	muestra-medica	\N	2021-03-10 04:25:34.740338	\N	t
18	ojo-seco	\N	2021-03-10 04:25:34.740338	\N	t
19	thealoz--alfa	\N	2021-03-10 04:25:34.740338	\N	t
20	thealoz-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
21	thealoz-beta	\N	2021-03-10 04:25:34.740338	\N	t
22	thealoz-duo-alfa	\N	2021-03-10 04:25:34.740338	\N	t
23	thealoz-duo-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
24	thealoz-duo-beta	\N	2021-03-10 04:25:34.740338	\N	t
25	thealoz-duo	\N	2021-03-10 04:25:34.740338	\N	t
26	thealoz	\N	2021-03-10 04:25:34.740338	\N	t
27	zaditen-alfa	\N	2021-03-10 04:25:34.740338	\N	t
28	zaditen-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
29	zaditen-beta	\N	2021-03-10 04:25:34.740338	\N	t
30	zaditen-gamma	\N	2021-03-10 04:25:34.740338	\N	t
31	zaditen	\N	2021-03-10 04:25:34.740338	\N	t
32	zyter-alfa	\N	2021-03-10 04:25:34.740338	\N	t
33	zyter-alfabeta	\N	2021-03-10 04:25:34.740338	\N	t
34	zyter-beta	\N	2021-03-10 04:25:34.740338	\N	t
35	zyter-gamma	\N	2021-03-10 04:25:34.740338	\N	t
36	zyter	\N	2021-03-10 04:25:34.740338	\N	t
\.


--
-- Data for Name: CTHEA_ROLE; Type: TABLE DATA; Schema: manager; Owner: postgres
--

COPY manager."CTHEA_ROLE" (id_role, name, code, creation_date, edit_date, status) FROM stdin;
1	GENERAL	GNL	2021-03-10 02:24:14.890841	\N	t
\.


--
-- Data for Name: THEA_DETAILS_ENDPOINT_AND_HOSTIN; Type: TABLE DATA; Schema: manager; Owner: postgres
--

COPY manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" (id_details_endpoint_and_hostin, id_details_hosting, id_endpoint, creation_date, edit_date, status) FROM stdin;
\.


--
-- Data for Name: THEA_DETAILS_HOSTING; Type: TABLE DATA; Schema: manager; Owner: postgres
--

COPY manager."THEA_DETAILS_HOSTING" (id_details_hosting, id_user, host, session_start, session_closure, type_log_out, creation_date, edit_date, status) FROM stdin;
\.


--
-- Data for Name: THEA_USER; Type: TABLE DATA; Schema: manager; Owner: postgres
--

COPY manager."THEA_USER" (id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, creation_date, edit_date, status) FROM stdin;
\.

