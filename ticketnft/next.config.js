/** @type {import('next').NextConfig} */
dns = require("dns");
dns.setDefaultResultOrder("ipv4first")
const nextConfig = {}

module.exports = nextConfig
