export default {
  items: [
    {
      name: "Redemption",
      url: "/dashboard",
      icon: "icon-present",
      badge: {
        variant: "warning",
        text: "beta"
      }
    },
    {
      title: true,
      name: "Redemption Codes",
      wrapper: {
        element: "", // required valid HTML5 element tag
        attributes: {} // optional valid JS object with JS API naming ex: { className: "my-class", style: { fontFamily: "Verdana" }, id: "my-id"}
      },
      class: "" // optional class names space delimited list for title item ex: "text-center"
    },
    {
      name: "All Codes",
      url: "/codes/all",
      icon: "icon-list"
    },
    {
      name: "Create Codes",
      url: "/codes/create",
      icon: "icon-pencil"
    },
    {
      title: true,
      name: "Management",
      wrapper: {
        element: "",
        attributes: {}
      }
    },
    {
      name: "Campaigns",
      url: "/",
      icon: "icon-flag",
      badge: {
        variant: "info",
        text: "Soon!"
      },
      attributes: {
        disabled: true
      }
    },
    {
      name: "Stats",
      url: "/",
      icon: "icon-pie-chart",
      badge: {
        variant: "info",
        text: "Soon!"
      },
      attributes: {
        disabled: true
      }
    }
  ]
};
