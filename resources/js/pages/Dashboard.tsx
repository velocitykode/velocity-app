import { AppLogo, Badge, IconButton, NavItem, StatCard, UserAvatar } from '@/components/ui';
import { useAppearance } from '@/hooks/use-appearance';
import { usePageTitle } from '@/hooks/use-page-title';
import { Link, usePage } from '@inertiajs/react';
import {
  Activity,
  ArrowRight,
  BookOpen,
  Check,
  Database,
  Gauge,
  LayoutGrid,
  LogOut,
  Moon,
  Route,
  Server,
  Settings,
  Shield,
  SquareMenu,
  Sun,
  Terminal,
  X,
  Zap,
} from 'lucide-react';
import { useState } from 'react';

const navigation = [
  { name: 'Dashboard', href: '/dashboard', icon: LayoutGrid, active: true },
  { name: 'Routes', href: '#', icon: Route },
  { name: 'Middleware', href: '#', icon: Shield },
  { name: 'Database', href: '#', icon: Database },
];

const resources = [
  { name: 'Documentation', href: 'https://github.com/velocitykode/velocity', icon: BookOpen, external: true },
  { name: 'Settings', href: '#', icon: Settings },
];

const stats = [
  { label: 'Routes', value: '12', icon: Route, color: 'blue' as const, change: '+2' },
  { label: 'Middleware', value: '6', icon: Shield, color: 'violet' as const },
  { label: 'Cache Hit', value: '94%', icon: Gauge, color: 'emerald' as const },
  { label: 'Uptime', value: '99.9%', icon: Server, color: 'amber' as const },
];

const tasks = [
  { task: 'Configure routes', done: true },
  { task: 'Set up middleware', done: true },
  { task: 'Connect database', done: false },
  { task: 'Configure cache', done: false },
  { task: 'Deploy to production', done: false },
];

const activities = [
  { action: 'Migration completed', time: '2m', icon: Database, color: 'emerald' as const },
  { action: 'Cache invalidated', time: '15m', icon: Server, color: 'blue' as const },
  { action: 'Route registered', time: '1h', icon: Route, color: 'violet' as const },
  { action: 'Middleware updated', time: '3h', icon: Shield, color: 'amber' as const },
  { action: 'Server restarted', time: '5h', icon: Terminal, color: 'rose' as const },
];

const activityColorClasses = {
  emerald: 'bg-emerald-100 text-emerald-600 dark:bg-emerald-900/50 dark:text-emerald-400',
  blue: 'bg-blue-100 text-[#1e3a8a] dark:bg-blue-900/50 dark:text-blue-400',
  violet: 'bg-violet-100 text-violet-600 dark:bg-violet-900/50 dark:text-violet-400',
  amber: 'bg-amber-100 text-amber-600 dark:bg-amber-900/50 dark:text-amber-400',
  rose: 'bg-rose-100 text-rose-600 dark:bg-rose-900/50 dark:text-rose-400',
};

export default function Dashboard() {
  usePageTitle('Dashboard');
  const { props } = usePage<{ auth: { user: { name: string; email: string } } }>();
  const { appearance, updateAppearance } = useAppearance();
  const user = props.auth?.user;
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const toggleTheme = () => {
    updateAppearance(appearance === 'dark' ? 'light' : 'dark');
  };

  const SidebarContent = () => (
    <>
      {/* Logo */}
      <div className="flex h-16 items-center justify-between border-b border-slate-100 px-4 dark:border-zinc-800">
        <AppLogo variant="responsive" size="xl" href="/" />
        <div className="flex items-center gap-2">
          <Badge variant="brand">v0.1</Badge>
          <IconButton
            icon={X}
            variant="default"
            aria-label="Close menu"
            onClick={() => setSidebarOpen(false)}
            className="lg:hidden"
          />
        </div>
      </div>

      {/* Nav */}
      <nav className="flex-1 overflow-y-auto p-4">
        <div className="mb-2 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400 dark:text-zinc-500">
          Platform
        </div>
        <div className="space-y-1">
          {navigation.map((item) => (
            <NavItem
              key={item.name}
              {...item}
              onClick={() => setSidebarOpen(false)}
            />
          ))}
        </div>

        <div className="mb-2 mt-8 px-3 text-xs font-semibold uppercase tracking-wider text-slate-400 dark:text-zinc-500">
          Resources
        </div>
        <div className="space-y-1">
          {resources.map((item) => (
            <NavItem key={item.name} {...item} />
          ))}
        </div>
      </nav>

      {/* User */}
      <div className="border-t border-slate-100 p-4 dark:border-zinc-800">
        <div className="flex items-center gap-3">
          <UserAvatar name={user?.name} />
          <div className="min-w-0 flex-1">
            <div className="truncate text-sm font-medium text-slate-900 dark:text-white">
              {user?.name || 'User'}
            </div>
            <div className="truncate text-xs text-slate-500 dark:text-zinc-400">
              {user?.email}
            </div>
          </div>
          <Link href="/logout" method="post" as="button">
            <IconButton
              icon={LogOut}
              variant="danger"
              aria-label="Log out"
              as="button"
              type="button"
            />
          </Link>
        </div>
      </div>
    </>
  );

  return (
    <div className="flex h-screen bg-slate-50 dark:bg-zinc-950">
      {/* Mobile sidebar overlay */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 z-40 bg-black/50 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}

      {/* Mobile sidebar */}
      <aside
        className={`fixed inset-y-0 left-0 z-50 w-72 transform bg-white transition-transform duration-200 ease-in-out lg:hidden dark:bg-zinc-900 ${
          sidebarOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <div className="flex h-full flex-col">
          <SidebarContent />
        </div>
      </aside>

      {/* Desktop sidebar */}
      <aside className="hidden w-72 flex-col border-r border-slate-200 bg-white lg:flex dark:border-zinc-800 dark:bg-zinc-900">
        <SidebarContent />
      </aside>

      {/* Main */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <header className="flex h-16 items-center justify-between border-b border-slate-200 bg-white px-4 lg:px-8 dark:border-zinc-800 dark:bg-zinc-900">
          <div className="flex items-center gap-3">
            <IconButton
              icon={SquareMenu}
              variant="ghost"
              size="lg"
              aria-label="Open menu"
              onClick={() => setSidebarOpen(true)}
              className="lg:hidden"
            />
            <h1 className="text-lg font-semibold text-slate-900 dark:text-white">Dashboard</h1>
          </div>
          <div className="flex items-center gap-2">
            <IconButton
              icon={appearance === 'dark' ? Sun : Moon}
              variant="default"
              aria-label="Toggle theme"
              onClick={toggleTheme}
            />
          </div>
        </header>

        {/* Content */}
        <main className="flex-1 overflow-y-auto p-4 lg:p-8">
          {/* Welcome */}
          <div className="mb-6 lg:mb-8">
            <Badge variant="brand" size="md" className="mb-2">
              <Zap className="h-4 w-4" />
              Ready to build
            </Badge>
            <h2 className="text-xl lg:text-2xl font-bold text-slate-900 dark:text-white">
              Welcome back, {user?.name?.split(' ')[0] || 'Developer'}
            </h2>
            <p className="mt-1 text-sm lg:text-base text-slate-500 dark:text-zinc-400">
              Your application overview
            </p>
          </div>

          {/* Stats */}
          <div className="mb-6 lg:mb-8 grid gap-4 grid-cols-2 lg:grid-cols-4">
            {stats.map((stat) => (
              <StatCard key={stat.label} {...stat} />
            ))}
          </div>

          {/* Two Column */}
          <div className="grid gap-4 lg:gap-6 lg:grid-cols-2">
            {/* Getting Started */}
            <div className="rounded-2xl border border-slate-200 bg-white dark:border-zinc-800 dark:bg-zinc-900">
              <div className="flex items-center justify-between border-b border-slate-100 p-4 lg:p-6 dark:border-zinc-800">
                <div className="flex items-center gap-3">
                  <div className="rounded-xl bg-blue-100 p-2 lg:p-2.5 text-[#1e3a8a] dark:bg-blue-900/50 dark:text-blue-400">
                    <Zap className="h-4 w-4 lg:h-5 lg:w-5" />
                  </div>
                  <div>
                    <h3 className="font-semibold text-slate-900 dark:text-white">Getting Started</h3>
                    <p className="text-xs lg:text-sm text-slate-500 dark:text-zinc-400">2 of 5 completed</p>
                  </div>
                </div>
                <span className="text-xl lg:text-2xl font-bold text-[#1e3a8a] dark:text-blue-400">40%</span>
              </div>
              <div className="divide-y divide-slate-100 dark:divide-zinc-800">
                {tasks.map((item, i) => (
                  <div key={i} className="flex items-center gap-3 p-3 lg:p-4 transition-colors hover:bg-slate-50 dark:hover:bg-zinc-800/50">
                    <div className={`flex h-5 w-5 lg:h-6 lg:w-6 items-center justify-center rounded-full ${
                      item.done ? 'bg-emerald-500 text-white' : 'border-2 border-slate-300 dark:border-zinc-600'
                    }`}>
                      {item.done && <Check className="h-3 w-3 lg:h-3.5 lg:w-3.5" />}
                    </div>
                    <span className={`flex-1 text-sm ${
                      item.done ? 'text-slate-400 line-through dark:text-zinc-500' : 'text-slate-700 dark:text-zinc-300'
                    }`}>
                      {item.task}
                    </span>
                    <ArrowRight className="h-4 w-4 text-slate-300 dark:text-zinc-600" />
                  </div>
                ))}
              </div>
              <div className="p-3 lg:p-4">
                <a
                  href="https://github.com/velocitykode/velocity"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="flex w-full items-center justify-center gap-2 rounded-xl bg-[#1e3a8a] py-2.5 lg:py-3 text-sm font-semibold text-white transition-colors hover:bg-[#1e3070]"
                >
                  <BookOpen className="h-4 w-4" />
                  View Documentation
                </a>
              </div>
            </div>

            {/* Activity */}
            <div className="rounded-2xl border border-slate-200 bg-white dark:border-zinc-800 dark:bg-zinc-900">
              <div className="flex items-center justify-between border-b border-slate-100 p-4 lg:p-6 dark:border-zinc-800">
                <div className="flex items-center gap-3">
                  <div className="rounded-xl bg-violet-100 p-2 lg:p-2.5 text-violet-600 dark:bg-violet-900/50 dark:text-violet-400">
                    <Activity className="h-4 w-4 lg:h-5 lg:w-5" />
                  </div>
                  <div>
                    <h3 className="font-semibold text-slate-900 dark:text-white">Recent Activity</h3>
                    <p className="text-xs lg:text-sm text-slate-500 dark:text-zinc-400">Live events</p>
                  </div>
                </div>
                <div className="flex items-center gap-2">
                  <span className="relative flex h-2 w-2 lg:h-2.5 lg:w-2.5">
                    <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75" />
                    <span className="relative inline-flex h-2 w-2 lg:h-2.5 lg:w-2.5 rounded-full bg-emerald-500" />
                  </span>
                  <span className="text-xs font-medium text-emerald-600 dark:text-emerald-400">Live</span>
                </div>
              </div>
              <div className="divide-y divide-slate-100 dark:divide-zinc-800">
                {activities.map((item, i) => (
                  <div key={i} className="flex items-center gap-3 p-3 lg:p-4 transition-colors hover:bg-slate-50 dark:hover:bg-zinc-800/50">
                    <div className={`rounded-lg p-1.5 lg:p-2 ${activityColorClasses[item.color]}`}>
                      <item.icon className="h-3.5 w-3.5 lg:h-4 lg:w-4" />
                    </div>
                    <span className="flex-1 text-sm text-slate-700 dark:text-zinc-300 truncate">{item.action}</span>
                    <span className="text-xs text-slate-400 dark:text-zinc-500">{item.time}</span>
                  </div>
                ))}
              </div>
              <div className="border-t border-slate-100 p-3 lg:p-4 dark:border-zinc-800">
                <button className="w-full text-center text-sm font-medium text-[#1e3a8a] hover:underline dark:text-blue-400">
                  View all activity →
                </button>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}
