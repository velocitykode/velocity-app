import { AppLogo, Badge } from '@/components/ui';
import { Link } from '@inertiajs/react';
import { type PropsWithChildren } from 'react';

interface AuthLayoutProps {
  title: string;
  description: string;
}

export default function AuthLayout({
  title,
  description,
  children,
}: PropsWithChildren<AuthLayoutProps>) {
  return (
    <div className="flex min-h-screen flex-col lg:flex-row">
      {/* Left - Brand (hidden on mobile, shown on lg+) */}
      <div className="relative hidden w-1/2 overflow-hidden bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 lg:flex lg:flex-col lg:justify-between lg:p-12">
        {/* Gradient orbs */}
        <div className="absolute -left-40 -top-40 h-80 w-80 rounded-full bg-blue-500/30 blur-3xl" />
        <div className="absolute -bottom-40 -right-40 h-80 w-80 rounded-full bg-[#dc2626]/15 blur-3xl" />
        <div className="absolute left-1/2 top-1/2 h-60 w-60 -translate-x-1/2 -translate-y-1/2 rounded-full bg-cyan-500/20 blur-3xl" />

        {/* Grid pattern */}
        <div className="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,0.02)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.02)_1px,transparent_1px)] bg-[size:60px_60px]" />

        {/* Content */}
        <div className="relative">
          <Link href="/">
            <img src="/logo-dark.png" alt="Velocity" className="h-28" />
          </Link>
        </div>

        <div className="relative">
          <Badge variant="default" dot dotColor="danger" className="mb-4 bg-white/10 text-white/70">
            Ready to Ship
          </Badge>
          <h2 className="text-4xl font-bold leading-tight text-white">
            Build faster with
            <br />
            <span className="bg-gradient-to-r from-[#1e3a8a] via-blue-400 to-slate-300 bg-clip-text text-transparent">
              pure Go power
            </span>
          </h2>
          <p className="mt-6 max-w-md text-lg text-slate-400">
            The modern web framework that combines Laravel&apos;s elegance with Go&apos;s raw performance.
          </p>

          {/* Features */}
          <div className="mt-10 grid grid-cols-2 gap-4">
            {[
              { value: '21', label: 'Packages' },
              { value: '0', label: 'Dependencies' },
              { value: '10x', label: 'Faster' },
              { value: '∞', label: 'Scale' },
            ].map((stat) => (
              <div key={stat.label} className="rounded-2xl bg-white/5 p-4">
                <div className="text-2xl font-bold text-white">{stat.value}</div>
                <div className="text-sm text-slate-400">{stat.label}</div>
              </div>
            ))}
          </div>
        </div>

        <div className="relative text-sm text-slate-500">
          © {new Date().getFullYear()} Velocity Framework
        </div>
      </div>

      {/* Right - Form */}
      <div className="flex w-full flex-1 flex-col bg-white dark:bg-zinc-950 lg:w-1/2">
        {/* Mobile header */}
        <div className="flex h-14 items-center justify-between border-b border-slate-100 px-4 lg:hidden dark:border-zinc-800">
          <AppLogo variant="icon" size="xl" href="/" />
          <Badge variant="brand">v0.1</Badge>
        </div>

        {/* Form */}
        <div className="flex flex-1 items-center justify-center p-6 lg:p-8">
          <div className="w-full max-w-sm">
            <div className="mb-6 lg:mb-8">
              <h1 className="text-xl lg:text-2xl font-bold text-slate-900 dark:text-white">{title}</h1>
              <p className="mt-2 text-sm lg:text-base text-slate-500 dark:text-slate-400">{description}</p>
            </div>

            {children}
          </div>
        </div>

        {/* Mobile footer */}
        <div className="border-t border-slate-100 p-4 text-center text-xs text-slate-400 lg:hidden dark:border-zinc-800 dark:text-slate-500">
          © {new Date().getFullYear()} Velocity Framework
        </div>
      </div>
    </div>
  );
}
