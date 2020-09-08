module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "sell", fields: ["item", "amount", "info", ] },
		{ type: "buy", fields: ["item", "amount", "info", ] },
		{ type: "asset", fields: ["title", "body", ] },
		{ type: "comment", fields: ["title", "body", ] },
		{ type: "post", fields: ["title", "body", ] },
  ],
};
